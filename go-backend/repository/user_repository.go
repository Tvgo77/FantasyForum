package repository

import (
	"context"
	"go-backend/domain"
	"go-backend/setup"
	"time"
)

type userRepository struct {
	database domain.Database
	env *setup.Env
}

func NewUserRepository(db domain.Database, env *setup.Env) domain.UserRepository {
	return &userRepository{database: db, env: env};
}

func (ur *userRepository) CheckExistByEmail(ctx context.Context, email string) (bool, error) {
	// Set timeout for database query
	ctx, cancel := context.WithTimeout(ctx, time.Duration(ur.env.TimeoutSeconds))
	defer cancel()

	count, err := ur.database.CountRows(ctx, &domain.User{Email: email})
	return count > 0, err
}

func (ur *userRepository) Create(ctx context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(ur.env.TimeoutSeconds))
	defer cancel()

	err := ur.database.InsertOne(ctx, user)
	return err
}