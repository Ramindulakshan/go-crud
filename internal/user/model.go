package user

import "github.com/google/uuid"

const (
	StatusActive   = "Active"
	StatusInactive = "Inactive"
)

type User struct {
	UserID    uuid.UUID `json:"userId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	Age       *int      `json:"age,omitempty"`
	Status    string    `json:"status"`
}

type CreateUserRequest struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"omitempty,min=7,max=25"`
	Age       *int   `json:"age" validate:"omitempty,gt=0"`
	Status    string `json:"status" validate:"omitempty,oneof=Active Inactive"`
}

type UpdateUserRequest struct {
	FirstName string `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string `json:"lastName" validate:"required,min=2,max=50"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"omitempty,min=7,max=25"`
	Age       *int   `json:"age" validate:"omitempty,gt=0"`
	Status    string `json:"status" validate:"omitempty,oneof=Active Inactive"`
}
