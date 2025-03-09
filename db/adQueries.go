package db

import (
	"context"

	"github.com/truongtu268/distributePriorityQueue/db/query"
)

type AdCreateQuery interface {
	CreateAd(ctx context.Context, arg query.CreateAdParams) error
}

type AdQueueQuery interface {
	UpdateAdStatus(ctx context.Context, arg query.UpdateAdStatusParams) error
}

type AdCronjobQuery interface {
	UpdateAdStatus(ctx context.Context, arg query.UpdateAdStatusParams) error
	GetAdByID(ctx context.Context, id string) (query.Ad, error)
	UpdateAdAnalysis(ctx context.Context, arg query.UpdateAdAnalysisParams) error
	UpdateAdRetry(ctx context.Context, arg query.UpdateAdRetryParams) error
}

type AdGetQuery interface {
	GetAdByID(ctx context.Context, id string) (query.Ad, error)
}

func NewAdCreateQuery(db query.DBTX) AdCreateQuery {
	return query.New(db)
}

func NewAdQueueQuery(db query.DBTX) AdQueueQuery {
	return query.New(db)
}

func NewAdCronjobQuery(db query.DBTX) AdCronjobQuery {
	return query.New(db)
}

func NewAdGetQuery(db query.DBTX) AdGetQuery {
	return query.New(db)
}
