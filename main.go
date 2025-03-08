package main

import (
	"context"
	"database/sql"
	"math/rand"
	"net/http"
	"time"

	"github.com/truongtu268/distributePriorityQueue/model"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var ctx = context.Background()

func main() {
	r := gin.Default()

	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // use your Redis server address
	})

	// Initialize PostgreSQL connection
	connStr := "user=youruser password=yourpassword dbname=yourdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	r.POST("/ads", func(c *gin.Context) {
		var newAd model.AdRequest
		if err := c.ShouldBindJSON(&newAd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate a unique adId (for simplicity, using a static value here)
		adId := "ad123" // In a real application, generate a unique ID

		// Set the status and createdAt
		status := "queued"
		createdAt := time.Now().Format(time.RFC3339)

		// Store the ad in the database using a raw SQL query
		query := `
			INSERT INTO ads (title, description, genre, target_audience, visual_elements, call_to_action, duration, priority, created_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`
		_, err := db.Exec(query, newAd.Title, newAd.Description, newAd.Genre, pq.Array(newAd.TargetAudience), pq.Array(newAd.VisualElements), newAd.CallToAction, newAd.Duration, newAd.Priority, createdAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store ad in database"})
			return
		}

		// Create the response object
		response := model.AdResponse{
			AdID:      adId,
			Status:    status,
			Priority:  newAd.Priority,
			CreatedAt: createdAt,
		}

		// Publish ad data to Redis stream
		err = rdb.XAdd(ctx, &redis.XAddArgs{
			Stream: "ads_stream",
			Values: map[string]interface{}{
				"adId":      adId,
				"title":     newAd.Title,
				"status":    status,
				"priority":  newAd.Priority,
				"createdAt": createdAt,
			},
		}).Err()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish to Redis stream"})
			return
		}

		// Respond with the specified format
		c.JSON(http.StatusOK, response)
	})

	r.GET("/ads/:id", func(c *gin.Context) {
		adId := c.Param("id")

		// Fetch ad details from the database
		var ad model.AdRequest
		var createdAt time.Time
		query := `
			SELECT title, description, genre, target_audience, visual_elements, call_to_action, duration, priority, created_at
			FROM ads
			WHERE id = $1
		`
		err := db.QueryRow(query, adId).Scan(&ad.Title, &ad.Description, &ad.Genre, pq.Array(&ad.TargetAudience), pq.Array(&ad.VisualElements), &ad.CallToAction, &ad.Duration, &ad.Priority, &createdAt)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, gin.H{"error": "Ad not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ad from database"})
			}
			return
		}

		// Seed the random number generator
		rand.Seed(time.Now().UnixNano())

		// Generate random analysis data
		effectivenessScore := rand.Float64() * 10 // Random score between 0 and 10
		strengths := []string{
			"Strong call to action with clear incentive",
			"Appeals to target audience's desire for progression",
		}
		improvementSuggestions := []string{
			"Consider adding social proof elements",
			"Highlight immediate gameplay satisfaction",
		}

		// Shuffle strengths and improvement suggestions
		rand.Shuffle(len(strengths), func(i, j int) { strengths[i], strengths[j] = strengths[j], strengths[i] })
		rand.Shuffle(len(improvementSuggestions), func(i, j int) {
			improvementSuggestions[i], improvementSuggestions[j] = improvementSuggestions[j], improvementSuggestions[i]
		})

		// Create the response object
		response := model.AdDetailResponse{
			AdID:   adId,
			Status: "completed",
			Analysis: model.AdAnalysis{
				EffectivenessScore:     effectivenessScore,
				Strengths:              strengths,
				ImprovementSuggestions: improvementSuggestions,
			},
			CreatedAt:   createdAt.Format(time.RFC3339),
			CompletedAt: time.Now().Format(time.RFC3339), // Simulate completion time
		}

		c.JSON(http.StatusOK, response)
	})

	r.PUT("/ads/:id/analysis", func(c *gin.Context) {
		adId := c.Param("id")

		var analysis model.AdAnalysis
		if err := c.ShouldBindJSON(&analysis); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Update the analysis data in the database
		query := `
			UPDATE ads
			SET effectiveness_score = $1, strengths = $2, improvement_suggestions = $3
			WHERE id = $4
		`
		_, err := db.Exec(query, analysis.EffectivenessScore, pq.Array(analysis.Strengths), pq.Array(analysis.ImprovementSuggestions), adId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ad analysis in database"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "Analysis updated successfully"})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
