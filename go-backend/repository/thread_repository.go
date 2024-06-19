package repository

import (
	"context"
	"go-backend/domain"
	"go-backend/setup"
	"time"

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
	ctx, cancel := context.WithTimeout(ctx, time.Duration(tr.env.TimeoutSeconds) * time.Second)
	defer cancel()

	err := tr.database.InsertOne(ctx, thread)
	return err
}

func (tr *threadRepository) Fetch(ctx context.Context, conds *domain.Thread) (*domain.Thread, error) {
	return nil, nil
}

func (tr *threadRepository) FetchBatch(ctx context.Context, conds *domain.Thread, batchSize int, batchIndex int) ([]domain.Thread, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(tr.env.TimeoutSeconds) * time.Second)
	defer cancel()

	start := batchSize * batchIndex
	var threads []domain.Thread
	err := tr.database.WithContext(ctx).Where(conds).Order("create_time").Limit(batchSize).Offset(start).Find(&threads).Error
	if err != nil {
		return nil, err
	}

	return threads, nil
}

func (tr *threadRepository) Update(ctx context.Context, old *domain.Thread, new *domain.Thread) error {
	return nil
}

func (tr *threadRepository) Increase(ctx context.Context, old *domain.Thread, column string, n int) error {
	return nil
}