package model

import "time"

// AdRequest represents the structure of an ad submission
type AdRequest struct {
	Title          string   `json:"title" binding:"required"`
	Description    string   `json:"description" binding:"required"`
	Genre          string   `json:"genre" binding:"required"`
	TargetAudience []string `json:"targetAudience" binding:"required"`
	VisualElements []string `json:"visualElements" binding:"required"`
	CallToAction   string   `json:"callToAction" binding:"required"`
	Duration       int      `json:"duration" binding:"required"`
	Priority       int      `json:"priority" binding:"required"`
}

// AdResponse represents the structure of the response for an ad submission
type AdResponse struct {
	AdID      string `json:"adId"`
	Status    string `json:"status"`
	Priority  int    `json:"priority"`
	CreatedAt string `json:"createdAt"`
}

// AdAnalysis represents the analysis details of an ad
type AdAnalysis struct {
	EffectivenessScore     float64  `json:"effectivenessScore"`
	Strengths              []string `json:"strengths"`
	ImprovementSuggestions []string `json:"improvementSuggestions"`
}

// AdDetailResponse represents the detailed response for an ad
type AdDetailResponse struct {
	AdID        string     `json:"adId"`
	Status      string     `json:"status"`
	Analysis    AdAnalysis `json:"analysis"`
	CreatedAt   time.Time  `json:"createdAt"`
	CompletedAt time.Time  `json:"completedAt"`
}

// PriorityQueueTask represents the task in priority queue
type PriorityQueueTask struct {
	AdID      string    `json:"adId"`
	CreatedAt time.Time `json:"createdAt"`
	Priority  int       `json:"priority"`
	RetryAt   time.Time `json:"retryAt"`
	RetryTime int       `json:"retryTime"`
}
