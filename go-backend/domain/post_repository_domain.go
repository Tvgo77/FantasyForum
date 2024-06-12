package domain

import (
	"context"
	"time"
)

type Post struct {
	ID uint  `json:"pid"`
	UserID uint   `json:"uid"`
	ThreadID uint  `json:"tid"`
	CreateTime time.Time  `json:"createTime"`
	Content string  `json:"content"`
	VoteCount string  `json:"voteCount"`
}

type PostRepository interface {
	Create(context.Context, *Post) error
	Fetch(ctx context.Context, conds *Post) (*Post, error)
	FetchBatch(ctx context.Context, conds *Post) ([]Post, error)
	Update(ctx context.Context, old *Post, new *Post) error
}