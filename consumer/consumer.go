package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // use your Redis server address
	})

	// Start consuming messages from the stream
	for {
		// Read messages from the stream
		streams, err := rdb.XRead(ctx, &redis.XReadArgs{
			Streams: []string{"ads_stream", "0"}, // "0" to read from the beginning
			Count:   10,                          // Number of messages to read at a time
			Block:   0,                           // Block indefinitely
		}).Result()

		if err != nil {
			log.Fatalf("Failed to read from Redis stream: %v", err)
		}

		// Process each message
		for _, stream := range streams {
			for _, message := range stream.Messages {
				fmt.Printf("Received message ID: %s\n", message.ID)
				for key, value := range message.Values {
					fmt.Printf("%s: %v\n", key, value)
				}
				fmt.Println()
			}
		}

		// Sleep for a short duration before the next read
		time.Sleep(1 * time.Second)
	}
}
