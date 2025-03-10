package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/truongtu268/distributePriorityQueue/model"
	"github.com/truongtu268/distributePriorityQueue/repo"
	"github.com/truongtu268/distributePriorityQueue/service/queue"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ctx              = context.Background()
	rdb              *redis.Client
	queueService     *queue.Service
	queueCronjobRepo repo.IAdCronjobRepo
)

func Execute(pgConnString string, redisHost string) {
	// Initialize Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr: redisHost, // use your Redis server address
	})

	queueService = queue.NewService(time.Minute * 10)
	pgPool, err := pgxpool.New(context.Background(), pgConnString)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return
	}
	queueCronjobRepo = repo.NewAdCronjobRepo(pgPool)
	// Start consuming messages
	go consumeMessages()
	go ProcessTaskInQueue()
	go ProcessTaskInQueue()
	go ProcessTaskInQueue()

	fmt.Println("Started consuming messages")
}

func consumeMessages() {
	for {
		// Read messages from the stream
		streams, err := rdb.XRead(ctx, &redis.XReadArgs{
			Streams: []string{"ads_stream", "0"}, // "0" to read from the beginning
			Count:   10,                          // Number of messages to read at a time
			Block:   0,                           // Block indefinitely
		}).Result()

		if err != nil {
			log.Printf("Failed to read from Redis stream: %v", err)
			continue
		}

		// Process each message
		for _, stream := range streams {
			for _, message := range stream.Messages {
				log.Printf("Received message ID: %s", message.ID)
				itemID, _ := message.Values["itemID"].(string)
				createdAt, _ := time.Parse(time.RFC3339, message.Values["createdAt"].(string))
				priority, _ := message.Values["priority"].(int)
				retryAt, _ := time.Parse(time.RFC3339, message.Values["retryAt"].(string))
				retryTime, _ := message.Values["retryTime"].(int)

				// Create a PriorityQueueTask
				task := model.PriorityQueueTask{
					ItemID:    itemID,
					CreatedAt: createdAt,
					Priority:  priority,
					RetryAt:   retryAt,
					RetryTime: retryTime,
				}
				// Enqueue the task
				queueService.Enqueue(task)
				err = queueCronjobRepo.InQueueTask(ctx, itemID)
				if err != nil {
					log.Printf("Failed to update task status: %v", err)
				}
			}
		}
		// Sleep for a short duration before the next read
		time.Sleep(1 * time.Second)
	}
}

func ProcessTaskInQueue() {
	for {
		task, ok := queueService.Dequeue()
		if !ok {
			time.Sleep(1 * time.Second)
			continue
		}
		ad, err := queueCronjobRepo.GetAdByID(ctx, task.ItemID)
		if err != nil {
			log.Printf("Failed to process task: %v", err)
			queueService.RetryQueue.Enqueue(task)
			continue
		}
		err = queueCronjobRepo.ProcessTask(ctx, task.ItemID)
		if err != nil {
			log.Printf("Failed to process task: %v", err)
			queueService.RetryQueue.Enqueue(task)
			queueCronjobRepo.RetryAd(ctx, ad.ID, ad.RetryTime.Int32+1)
			continue
		}
		adAnalysis, err := CallExternalAnalysis()
		if err != nil {
			log.Printf("Failed to call external analysis: %v", err)
			queueService.RetryQueue.Enqueue(task)
			queueCronjobRepo.RetryAd(ctx, ad.ID, ad.RetryTime.Int32+1)
			continue
		}
		err = queueCronjobRepo.AddAdAnalysis(ctx, task.ItemID, adAnalysis)
		if err != nil {
			log.Printf("Failed to add ad analysis: %v", err)
			queueService.RetryQueue.Enqueue(task)
			queueCronjobRepo.RetryAd(ctx, ad.ID, ad.RetryTime.Int32+1)
			continue
		}
	}
}

// CallExternalAnalysis calls the external analysis API and returns the ad analysis
func CallExternalAnalysis() (model.AdAnalysis, error) {
	response, err := http.Get("http://localhost:8000/ad-analysis")
	if err != nil {
		log.Printf("Failed to call external analysis: %v", err)
		return model.AdAnalysis{}, err
	}
	defer response.Body.Close()

	var adAnalysis model.AdAnalysis
	err = json.NewDecoder(response.Body).Decode(&adAnalysis)
	if err != nil {
		log.Printf("Failed to decode external analysis response: %v", err)
		return model.AdAnalysis{}, err
	}
	return adAnalysis, nil
}
