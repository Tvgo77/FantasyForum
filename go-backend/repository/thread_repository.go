package repository

import (
	"context"
	"go-backend/domain"
	"go-backend/setup"

	"github.com/redis/go-redis/v9"
)


type threadRepository struct {
	database domain.Database
	cache *redis.Client
	env *setup.Env
}

func NewThreadRepository(db domain.Database, rdb *redis.Client, env *setup.Env) domain.ThreadRepository {
	return &threadRepository{database: db, cache: rdb, env: env}
}

func (tr *threadRepository) Create(ctx context.Context, thread *domain.Thread) error {
	return nil
}

func (tr *threadRepository) Fetch(ctx context.Context, conds *domain.Thread) (*domain.Thread, error) {
	return nil, nil
}

func (tr *threadRepository) FetchBatch(ctx context.Context, conds *domain.Thread) ([]domain.Thread, error) {
	return nil, nil
}

func (tr *threadRepository) Update(ctx context.Context, old *domain.Thread, new *domain.Thread) error {
	return nil
}