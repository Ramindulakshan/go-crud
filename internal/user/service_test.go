package user

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

type stubRepo struct {
	createFn func(context.Context, CreateUserRequest) (User, error)
	getFn    func(context.Context, uuid.UUID) (User, error)
	listFn   func(context.Context) ([]User, error)
	updateFn func(context.Context, uuid.UUID, UpdateUserRequest) (User, error)
	deleteFn func(context.Context, uuid.UUID) error
}

func (s stubRepo) Create(ctx context.Context, input CreateUserRequest) (User, error) {
	if s.createFn != nil {
		return s.createFn(ctx, input)
	}
	return User{}, nil
}

func (s stubRepo) GetByID(ctx context.Context, id uuid.UUID) (User, error) {
	if s.getFn != nil {
		return s.getFn(ctx, id)
	}
	return User{}, nil
}

func (s stubRepo) List(ctx context.Context) ([]User, error) {
	if s.listFn != nil {
		return s.listFn(ctx)
	}
	return nil, nil
}

func (s stubRepo) Update(ctx context.Context, id uuid.UUID, input UpdateUserRequest) (User, error) {
	if s.updateFn != nil {
		return s.updateFn(ctx, id, input)
	}
	return User{}, nil
}

func (s stubRepo) Delete(ctx context.Context, id uuid.UUID) error {
	if s.deleteFn != nil {
		return s.deleteFn(ctx, id)
	}
	return nil
}

func TestServiceCreateRejectsInvalidPhone(t *testing.T) {
	svc := NewService(stubRepo{})

	_, err := svc.Create(context.Background(), CreateUserRequest{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
		Phone:     "0712345678",
	})
	if err == nil {
		t.Fatal("expected validation error for non-E.164 phone")
	}
}

func TestServiceUpdateRequiresAtLeastOneField(t *testing.T) {
	svc := NewService(stubRepo{})

	_, err := svc.Update(context.Background(), uuid.New(), UpdateUserRequest{})
	if err == nil {
		t.Fatal("expected error for empty patch payload")
	}
}

func TestServiceUpdateCallsRepositoryForValidPatch(t *testing.T) {
	called := false
	firstName := "Jane"
	id := uuid.New()
	svc := NewService(stubRepo{
		updateFn: func(_ context.Context, gotID uuid.UUID, input UpdateUserRequest) (User, error) {
			called = true
			if gotID != id {
				t.Fatalf("unexpected id: %v", gotID)
			}
			if input.FirstName == nil || *input.FirstName != "Jane" {
				t.Fatalf("unexpected patch payload: %+v", input)
			}
			return User{UserID: id, FirstName: "Jane"}, nil
		},
	})

	_, err := svc.Update(context.Background(), id, UpdateUserRequest{FirstName: &firstName})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !called {
		t.Fatal("expected repository update to be called")
	}
}
