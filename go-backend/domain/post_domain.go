package domain

import "context"

type PostUsecase interface {
	GetPostsByThreadAndPage(ctx context.Context, tid uint, page uint) ([]Post, error)
	CreatePost(ctx context.Context, post *Post) error
}
