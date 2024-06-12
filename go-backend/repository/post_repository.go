package repository

import (
	"context"
	"go-backend/domain"
	"go-backend/setup"

	"github.com/redis/go-redis/v9"
)

type postRepository struct {
	database domain.Database
	cache *redis.Client
	env *setup.Env
}

func NewPostRepository(db domain.Database, rdb *redis.Client, env *setup.Env) domain.PostRepository {
	return &postRepository{database: db, cache: rdb, env: env}
}

func (pr *postRepository) Create(context.Context, *domain.Post) error {
	return nil
}

func (pr *postRepository) FetchBatch(ctx context.Context, conds *domain.Post) ([]domain.Post, error) {
	return nil, nil
}

func (pr *postRepository) Update(ctx context.Context, old *domain.Post, new *domain.Post) error {
	return nil
}