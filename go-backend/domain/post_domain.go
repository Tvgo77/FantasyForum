package domain

// GET
type FetchPostsRequest struct {

}

type FetchPostsResponse struct {
	Posts []Post  `json:"posts"`
}

// Post
type CreatePostRequest struct {
	Post Post  `json:"post"`
}

type CreatePostResponse struct {

}

type PostUsecase interface {
	
}
