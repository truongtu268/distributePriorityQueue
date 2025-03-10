package external

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// AdAnalysis represents the analysis details of an ad
type AdAnalysis struct {
	EffectivenessScore     float64  `json:"effectivenessScore"`
	Strengths              []string `json:"strengths"`
	ImprovementSuggestions []string `json:"improvementSuggestions"`
}

// GetAdAnalysis simulates fetching ad analysis data with a 3-second delay
func GetAdAnalysis() AdAnalysis {
	time.Sleep(3 * time.Second) // Simulate delay

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random effectiveness score between 0 and 100
	effectivenessScore := rand.Float64() * 100

	return AdAnalysis{
		EffectivenessScore:     effectivenessScore,
		Strengths:              []string{"Engaging content", "Strong call to action"},
		ImprovementSuggestions: []string{"Improve targeting", "Enhance visuals"},
	}
}

// AdAnalysisHandler handles HTTP requests for ad analysis
func AdAnalysisHandler(w http.ResponseWriter, r *http.Request) {
	analysis := GetAdAnalysis()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(analysis)
}

func Run() {
	http.HandleFunc("/ad-analysis", AdAnalysisHandler)

	// Start the server on port 8000
	log.Println("Server starting on port 8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
