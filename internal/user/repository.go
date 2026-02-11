package user

import (
	"context"
	"database/sql"
	"errors"

	db "go-crud/internal/db/sqlc"

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

type PostgresRepository struct {
	q *db.Queries
}

func NewPostgresRepository(q *db.Queries) *PostgresRepository {
	return &PostgresRepository{q: q}
}

func (r *PostgresRepository) Create(ctx context.Context, input CreateUserRequest) (User, error) {
	row, err := r.q.CreateUser(ctx, db.CreateUserParams{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     toNullString(input.Phone),
		Age:       toNullInt32(input.Age),
		Status:    resolveStatus(input.Status),
	})
	if err != nil {
		return User{}, err
	}
	return fromDBUser(row), nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id uuid.UUID) (User, error) {
	row, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrNotFound
		}
		return User{}, err
	}
	return fromDBUser(row), nil
}

func (r *PostgresRepository) List(ctx context.Context) ([]User, error) {
	rows, err := r.q.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	users := make([]User, 0, len(rows))
	for _, row := range rows {
		users = append(users, fromDBUser(row))
	}
	return users, nil
}

func (r *PostgresRepository) Update(ctx context.Context, id uuid.UUID, input UpdateUserRequest) (User, error) {
	existing, err := r.GetByID(ctx, id)
	if err != nil {
		return User{}, err
	}

	firstName := existing.FirstName
	if input.FirstName != nil {
		firstName = *input.FirstName
	}

	lastName := existing.LastName
	if input.LastName != nil {
		lastName = *input.LastName
	}

	email := existing.Email
	if input.Email != nil {
		email = *input.Email
	}

	phone := existing.Phone
	if input.Phone != nil {
		phone = *input.Phone
	}

	age := existing.Age
	if input.Age != nil {
		v := *input.Age
		age = &v
	}

	status := existing.Status
	if input.Status != nil {
		status = *input.Status
	}

	row, err := r.q.UpdateUser(ctx, db.UpdateUserParams{
		UserID:    id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     toNullString(phone),
		Age:       toNullInt32(age),
		Status:    resolveStatus(status),
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrNotFound
		}
		return User{}, err
	}
	return fromDBUser(row), nil
}

func (r *PostgresRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if _, err := r.q.GetUserByID(ctx, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}
		return err
	}
	return r.q.DeleteUser(ctx, id)
}

func fromDBUser(u db.User) User {
	var age *int
	if u.Age.Valid {
		v := int(u.Age.Int32)
		age = &v
	}

	phone := ""
	if u.Phone.Valid {
		phone = u.Phone.String
	}

	return User{
		UserID:    u.UserID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Phone:     phone,
		Age:       age,
		Status:    u.Status,
	}
}

func toNullString(v string) sql.NullString {
	if v == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: v, Valid: true}
}

func toNullInt32(v *int) sql.NullInt32 {
	if v == nil {
		return sql.NullInt32{}
	}
	return sql.NullInt32{Int32: int32(*v), Valid: true}
}

func resolveStatus(status string) string {
	if status == "" {
		return StatusActive
	}
	return status
}
