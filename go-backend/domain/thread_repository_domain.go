package domain

import (
	"context"
	"time"
)

type Thread struct {
	ID uint  `json:"tid"`
	UserID uint  `json:"uid"`
	Title string  `json:"title"`
	CreateTime time.Time  `json:"createTime"`
	LastPostTime time.Time  `json:"lastPostTime"`
	ViewCount int  `json:"viewCount"`
	PostCount int  `json:"postCount"`
	VoteCount int  `json:"voteCount"`
}

type ThreadRepository interface {
	Create(context.Context, *Thread) error
	Fetch(ctx context.Context, conds *Thread) (*Thread, error)
	FetchBatch(ctx context.Context, conds *Thread) ([]Thread, error)
	Update(ctx context.Context, old *Thread, new *Thread) error
}