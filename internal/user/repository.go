package user

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"
)

var ErrNotFound = errors.New("user not found")

type Repository interface {
	Create(ctx context.Context, input CreateUserRequest) (User, error)
	GetByID(ctx context.Context, id uuid.UUID) (User, error)
	List(ctx context.Context) ([]User, error)
	Update(ctx context.Context, id uuid.UUID, input UpdateUserRequest) (User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type InMemoryRepository struct {
	mu    sync.RWMutex
	users map[uuid.UUID]User
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{users: make(map[uuid.UUID]User)}
}

func (r *InMemoryRepository) Create(_ context.Context, input CreateUserRequest) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	status := input.Status
	if status == "" {
		status = StatusActive
	}

	u := User{
		UserID:    uuid.New(),
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		Age:       input.Age,
		Status:    status,
	}
	r.users[u.UserID] = u
	return u, nil
}

func (r *InMemoryRepository) GetByID(_ context.Context, id uuid.UUID) (User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	u, ok := r.users[id]
	if !ok {
		return User{}, ErrNotFound
	}
	return u, nil
}

func (r *InMemoryRepository) List(_ context.Context) ([]User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	out := make([]User, 0, len(r.users))
	for _, u := range r.users {
		out = append(out, u)
	}
	return out, nil
}

func (r *InMemoryRepository) Update(_ context.Context, id uuid.UUID, input UpdateUserRequest) (User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, ok := r.users[id]
	if !ok {
		return User{}, ErrNotFound
	}

	status := input.Status
	if status == "" {
		status = StatusActive
	}

	u := User{
		UserID:    id,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		Age:       input.Age,
		Status:    status,
	}
	r.users[id] = u
	return u, nil
}

func (r *InMemoryRepository) Delete(_ context.Context, id uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.users[id]; !ok {
		return ErrNotFound
	}
	delete(r.users, id)
	return nil
}
