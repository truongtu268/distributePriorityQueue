package repo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/truongtu268/distributePriorityQueue/db"
	"github.com/truongtu268/distributePriorityQueue/db/query"
	"github.com/truongtu268/distributePriorityQueue/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type ICreateAdRepo interface {
	CreateAd(ctx context.Context, dto model.AdRequest) (model.AdResponse, error)
}

type CreateAdRepo struct {
	queries db.AdCreateQuery
}

func (r *CreateAdRepo) CreateAd(ctx context.Context, dto model.AdRequest) (model.AdResponse, error) {
	adID := uuid.New()
	arg := query.CreateAdParams{
		ID: adID.String(),
		Title: pgtype.Text{
			String: dto.Title,
			Valid:  true,
		},
		Description: pgtype.Text{
			String: dto.Description,
			Valid:  true,
		},
		Status: pgtype.Text{
			String: string(model.SubmittedStatus),
			Valid:  true,
		},
		Genre: pgtype.Text{
			String: dto.Genre,
			Valid:  true,
		},
		TargetAudiences: dto.TargetAudience,
		VisualElements:  dto.VisualElements,
		CallToAction: pgtype.Text{
			String: dto.CallToAction,
			Valid:  true,
		},
		Duration:  pgtype.Int4{Int32: int32(dto.Duration), Valid: true},
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	if dto.Priority != 0 {
		arg.Priority = pgtype.Int4{Int32: int32(dto.Priority), Valid: true}
	} else {
		arg.Priority = pgtype.Int4{Int32: 1, Valid: true}
	}
	err := r.queries.CreateAd(ctx, arg)
	if err != nil {
		return model.AdResponse{}, err
	}

	return model.AdResponse{
		AdID:      adID.String(),
		Status:    arg.Status.String,
		CreatedAt: arg.CreatedAt.Time,
		Priority:  int(arg.Priority.Int32),
	}, nil
}

func NewCreateAdRepo(q query.DBTX) *CreateAdRepo {
	return &CreateAdRepo{
		queries: db.NewAdCreateQuery(q),
	}
}

type IAdQueueRepo interface {
	UpdateAdStatus(ctx context.Context, adID string) error
}

type AdQueueRepo struct {
	queries db.AdQueueQuery
}

func (r *AdQueueRepo) UpdateAdStatus(ctx context.Context, adID string) error {
	arg := query.UpdateAdStatusParams{
		ID: adID,
		Status: pgtype.Text{
			String: string(model.QueuedStatus),
			Valid:  true,
		},
	}
	return r.queries.UpdateAdStatus(ctx, arg)
}

func NewAdQueueRepo(q query.DBTX) *AdQueueRepo {
	return &AdQueueRepo{
		queries: db.NewAdQueueQuery(q),
	}
}

type IAdCronjobRepo interface {
	ProcessTask(ctx context.Context, adID string) error
	InQueueTask(ctx context.Context, adID string) error
	AddAdAnalysis(ctx context.Context, adID string, analysis model.AdAnalysis) error
	GetAdByID(ctx context.Context, adID string) (query.Ad, error)
	RetryAd(ctx context.Context, adID string, retryTime int32) error
}

type AdCronjobRepo struct {
	queries db.AdCronjobQuery
}

func (r *AdCronjobRepo) ProcessTask(ctx context.Context, adID string) error {
	arg := query.UpdateAdStatusParams{
		ID: adID,
		Status: pgtype.Text{
			String: string(model.ProcessingStatus),
			Valid:  true,
		},
	}
	return r.queries.UpdateAdStatus(ctx, arg)
}

func (r *AdCronjobRepo) InQueueTask(ctx context.Context, adID string) error {
	arg := query.UpdateAdStatusParams{
		ID: adID,
		Status: pgtype.Text{
			String: string(model.QueuedStatus),
			Valid:  true,
		},
	}
	return r.queries.UpdateAdStatus(ctx, arg)
}

func (r *AdCronjobRepo) AddAdAnalysis(ctx context.Context, adID string, analysis model.AdAnalysis) error {
	by, err := json.Marshal(analysis)
	if err != nil {
		return errors.New(fmt.Sprintf("error marshaling analysis: %v", err))
	}

	arg := query.UpdateAdAnalysisParams{
		ID: adID,
		Status: pgtype.Text{
			String: string(model.CompletedStatus),
			Valid:  true,
		},
		Analysis:    by,
		CompletedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}
	return r.queries.UpdateAdAnalysis(ctx, arg)
}

func (r *AdCronjobRepo) GetAdByID(ctx context.Context, adID string) (query.Ad, error) {
	ad, err := r.queries.GetAdByID(ctx, adID)
	if err != nil {
		return query.Ad{}, errors.New(fmt.Sprintf("error getting ad: %v", err))
	}
	return ad, nil
}

func (r *AdCronjobRepo) RetryAd(ctx context.Context, adID string, retryTime int32) error {
	arg := query.UpdateAdRetryParams{
		ID:        adID,
		RetriedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		RetryTime: pgtype.Int4{Int32: retryTime, Valid: true},
	}
	return r.queries.UpdateAdRetry(ctx, arg)
}

func NewAdCronjobRepo(q query.DBTX) *AdCronjobRepo {
	return &AdCronjobRepo{
		queries: db.NewAdCronjobQuery(q),
	}
}

type IGetAdRepo interface {
	GetAdByID(ctx context.Context, adID string) (query.Ad, error)
}

type GetAdRepo struct {
	queries db.AdGetQuery
}

func (r *GetAdRepo) GetAdByID(ctx context.Context, adID string) (query.Ad, error) {
	ad, err := r.queries.GetAdByID(ctx, adID)
	if err != nil {
		return query.Ad{}, errors.New(fmt.Sprintf("error getting ad: %v", err))
	}
	return ad, nil
}

func NewGetAdRepo(q query.DBTX) *GetAdRepo {
	return &GetAdRepo{
		queries: db.NewAdGetQuery(q),
	}
}
