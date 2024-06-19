package usecase

import (
	"context"
	"go-backend/domain"
	"go-backend/setup"
)

type postUsecase struct {
	postRepository domain.PostRepository
	env *setup.Env
}

func NewPostUsecase(pr domain.PostRepository, env *setup.Env) domain.PostUsecase {
	return &postUsecase{postRepository: pr, env: env}
}

func (pu *postUsecase) GetPostsByThreadAndPage(ctx context.Context, tid uint, page uint) ([]domain.Post, error) {
	// Check cache first
	conds := &domain.Post{ThreadID: tid}
	posts, err := pu.postRepository.FetchBatch(ctx, conds, pu.env.PostsPerPage, int(page) - 1)
	return posts, err
}

func (pu *postUsecase) CreatePost(ctx context.Context, post *domain.Post) error {
	err := pu.postRepository.Create(ctx, post)
	return err
}