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
	conds := &domain.Thread{ID: tid}
	thread, err := tu.threadRepository.Fetch(ctx, conds)
	if err != nil {
		return nil, err
	}
	return thread, nil
}

func (tu *threadUsecase) UpdateViewCount(ctx context.Context, tid uint, increase int) error {
	old := &domain.Thread{ID: tid}
	err := tu.threadRepository.Increase(ctx, old, "viewCount", 1)
	return err
}

func (tu *threadUsecase) GetThreadsByPage(ctx context.Context, page uint, order string) ([]domain.Thread, error) {
	conds := &domain.Thread{}
	threads, err := tu.threadRepository.FetchBatch(ctx, conds, tu.env.ThreadsPerPage, int(page) - 1)
	return threads, err 
}

func (tu *threadUsecase) CreateThread(ctx context.Context, thread *domain.Thread) error {
	err := tu.threadRepository.Create(ctx, thread)
	return err
}