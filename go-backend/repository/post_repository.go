package repository

import (
	"context"
	"go-backend/domain"
	"go-backend/setup"
)

type postRepository struct {
	database domain.Database
	env *setup.Env
}

func NewPostRepository(db domain.Database, env *setup.Env) domain.PostRepository {
	return &postRepository{database: db, env: env}
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