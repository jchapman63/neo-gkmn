// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"context"
)

type Querier interface {
	ListMonsters(ctx context.Context) ([]Monster, error)
}

var _ Querier = (*Queries)(nil)
