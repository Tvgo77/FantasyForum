package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"go-backend/domain"
	"go-backend/setup"
	"time"

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

func (pr *postRepository) Create(ctx context.Context, post *domain.Post) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(pr.env.TimeoutSeconds) * time.Second)
	defer cancel()

	err := pr.database.InsertOne(ctx, post)
	return err
}

func (pr *postRepository) FetchBatch(ctx context.Context, conds *domain.Post, batchSize int, batchIndex int) ([]domain.Post, error) {
	// Check cache first
	ctx, cancel := context.WithTimeout(ctx, time.Duration(pr.env.TimeoutSeconds) * time.Second)
	defer cancel()

	key := fmt.Sprintf("page:%d", conds.ThreadID)
	n, err :=pr.cache.Exists(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	start := batchSize * batchIndex
	stop := start + batchSize

	if n > 0 {
		// Retrive posts from cache
		path := fmt.Sprintf("$[%d:%d]", start, stop)
		result, err := pr.cache.JSONGet(ctx, key, path).Result()
		if err != nil {
			return nil, err
		}

		var posts []domain.Post
		err = json.Unmarshal([]byte(result), &posts)
		if err != nil {
			return nil, err
		}

		return posts, nil
	}

	// If not in cache Fetch from sql database
	var posts []domain.Post
	err = pr.database.WithContext(ctx).Where(conds).Order("create_time").Limit(batchSize).Offset(start).Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pr *postRepository) Update(ctx context.Context, old *domain.Post, new *domain.Post) error {
	return nil
}