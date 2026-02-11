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
	Phone     string `json:"phone" validate:"omitempty,e164"`
	Age       *int   `json:"age" validate:"omitempty,gt=0"`
	Status    string `json:"status" validate:"omitempty,oneof=Active Inactive"`
}

type UpdateUserRequest struct {
	FirstName *string `json:"firstName" validate:"omitempty,min=2,max=50"`
	LastName  *string `json:"lastName" validate:"omitempty,min=2,max=50"`
	Email     *string `json:"email" validate:"omitempty,email"`
	Phone     *string `json:"phone" validate:"omitempty,e164"`
	Age       *int    `json:"age" validate:"omitempty,gt=0"`
	Status    *string `json:"status" validate:"omitempty,oneof=Active Inactive"`
}

func (u UpdateUserRequest) HasUpdates() bool {
	return u.FirstName != nil || u.LastName != nil || u.Email != nil || u.Phone != nil || u.Age != nil || u.Status != nil
}
