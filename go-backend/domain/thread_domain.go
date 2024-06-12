package domain

import "context"

// GET
type FetchThreadsRequest struct {

}

type FetchThreadsResponse struct {
	Threads []Thread  `json:"threads"`
}

type FetchOneThreadRequest struct {

}

type FetchOneThreadResponse struct {
	Thread Thread  `json:"thread"`
	Posts []Post  `json:"posts"`
}

// POST
type CreateThreadRequest struct {
	Thread Thread `json:"thread" bind:"required"`
	Post Post `json:"post" bind:"required"`
}

type CreateThreadResponse struct {

}

type CreatePostRequest struct {
	Post Post  `json:"post" bind:"required"`
}

type CreatePostResponse struct {
	
}

type ThreadUsecase interface {
	GetThreadByID(ctx context.Context, tid uint) (*Thread, error)
	UpdateViewCount(ctx context.Context, tid uint, increase int) error
	GetThreadsByPage(ctx context.Context, page uint, order string) ([]Thread, error)
	CreateThread(ctx context.Context, thead *Thread) error
}