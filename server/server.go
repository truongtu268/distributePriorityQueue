package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/truongtu268/distributePriorityQueue/model"
	"github.com/truongtu268/distributePriorityQueue/repo"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AdServer struct {
	pgPool       *pgxpool.Pool
	createAdRepo repo.ICreateAdRepo
	getAdRepo    repo.IGetAdRepo
	redisClient  *redis.Client
}

func (s *AdServer) CreateAd(c *gin.Context) {
	var newAd model.AdRequest
	if err := c.ShouldBindJSON(&newAd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := s.createAdRepo.CreateAd(c.Request.Context(), newAd)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store ad in database"})
		return
	}

	// Publish ad data to Redis stream
	err = s.redisClient.XAdd(c.Request.Context(), &redis.XAddArgs{
		Stream: "ads_stream",
		Values: map[string]interface{}{
			"adId":      res.AdID,
			"title":     newAd.Title,
			"status":    res.Status,
			"priority":  res.Priority,
			"createdAt": res.CreatedAt,
		},
	}).Err()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to publish to Redis stream"})
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (s *AdServer) GetAd(c *gin.Context) {
	adId := c.Param("id")

	// Fetch ad details from the database using sqlc generated code
	ad, err := s.getAdRepo.GetAdByID(c.Request.Context(), adId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Ad not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ad from database"})
		}
		return
	}
	analysis := model.AdAnalysis{}
	err = json.Unmarshal(ad.Analysis, &analysis)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal ad analysis"})
		return
	}
	res := model.AdDetailResponse{
		AdID:        adId,
		Status:      ad.Status.String,
		CreatedAt:   ad.CreatedAt.Time,
		CompletedAt: ad.CompletedAt.Time,
		Analysis:    analysis,
	}
	c.JSON(http.StatusOK, res)
}

func (s *AdServer) Run() {
	r := gin.Default()
	r.POST("/ads", s.CreateAd)
	r.GET("/ads/:id", s.GetAd)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func NewAdServer(pgConnString string, options redis.Options) (*AdServer, error) {
	// Initialize PostgreSQL connection pool
	pgPool, err := pgxpool.New(context.Background(), pgConnString)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}
	return &AdServer{
		pgPool:       pgPool,
		redisClient:  redis.NewClient(&options),
		createAdRepo: repo.NewCreateAdRepo(pgPool),
		getAdRepo:    repo.NewGetAdRepo(pgPool),
	}, nil
}

func (s *AdServer) Close() {
	if s.pgPool != nil {
		s.pgPool.Close()
	}
}
