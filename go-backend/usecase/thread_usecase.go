package usecase

import (
	"context"
	"go-backend/domain"
	"go-backend/setup"
)

type threadUsecase struct {
	threadRepository domain.ThreadRepository
	env *setup.Env
}

func NewThreadUsecase(tr domain.ThreadRepository, env *setup.Env) domain.ThreadUsecase {
	return &threadUsecase{threadRepository: tr, env: env}
}

func (tu *threadUsecase) GetThreadByID(ctx context.Context, tid uint) (*domain.Thread, error) {
	return nil, nil
}

func (tu *threadUsecase) UpdateViewCount(ctx context.Context, tid uint, increase int) error {
	return nil
}

func (tu *threadUsecase) GetThreadsByPage(ctx context.Context, page uint, order string) ([]domain.Thread, error) {
	return nil, nil 
}

func (tu *threadUsecase) CreateThread(ctx context.Context, thead *domain.Thread) error {
	return nil
}