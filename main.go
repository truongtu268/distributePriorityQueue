package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ad represents the structure of an ad submission
type Ad struct {
	Title          string   `json:"title" binding:"required"`
	Description    string   `json:"description" binding:"required"`
	Genre          string   `json:"genre" binding:"required"`
	TargetAudience []string `json:"targetAudience" binding:"required"`
	VisualElements []string `json:"visualElements" binding:"required"`
	CallToAction   string   `json:"callToAction" binding:"required"`
	Duration       int      `json:"duration" binding:"required"`
	Priority       int      `json:"priority" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/ads", func(c *gin.Context) {
		var newAd Ad
		if err := c.ShouldBindJSON(&newAd); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Here you would typically save the ad to a database
		// For now, we'll just return the ad as a confirmation
		c.JSON(http.StatusOK, gin.H{"status": "Ad created successfully", "ad": newAd})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
