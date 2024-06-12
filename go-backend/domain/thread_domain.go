package domain

// GET
type FetchThreadsRequest struct {

}

type FetchThreadsResponse struct {
	Thread Thread  `json:"thread"`
}

// POST
type CreateThreadRequest struct {
	Thread Thread `json:"thread" bind:"required"`
	Post Post `json:"post" bind:"required"`
}

type CreateThreadResponse struct {

}

type ThreadUsecase interface {

}