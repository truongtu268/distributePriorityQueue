// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package query

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Ad struct {
	ID              string
	Title           pgtype.Text
	Description     pgtype.Text
	Status          pgtype.Text
	Genre           pgtype.Text
	TargetAudiences []string
	VisualElements  []string
	Analysis        []byte
	CallToAction    pgtype.Text
	Duration        pgtype.Int4
	Priority        pgtype.Int4
	CreatedAt       pgtype.Timestamptz
	RetriedAt       pgtype.Timestamptz
	CompletedAt     pgtype.Timestamptz
	RetryTime       pgtype.Int4
}
