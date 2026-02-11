package user

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Service struct {
	repo     Repository
	validate *validator.Validate
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo, validate: validator.New()}
}

func (s *Service) Create(ctx context.Context, input CreateUserRequest) (User, error) {
	if err := s.validate.Struct(input); err != nil {
		return User{}, err
	}
	return s.repo.Create(ctx, input)
}

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]User, error) {
	return s.repo.List(ctx)
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, input UpdateUserRequest) (User, error) {
	if err := s.validate.Struct(input); err != nil {
		return User{}, err
	}
	return s.repo.Update(ctx, id, input)
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
