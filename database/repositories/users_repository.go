/*
Copyright 2025 Alexis Bouchez <alexbcz@proton.me>

This file is part of Vauban.

Vauban is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

Vauban is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with Vauban. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing options, please contact Alexis Bouchez at alexbcz@proton.me
*/
package repositories

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/alexisbcz/vauban/database/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository interface {
	// Store creates a user record in the database, mutating fields generated by the DBMS.
	Store(ctx context.Context, user *models.User) error

	// GetByID retrieves a user by their ID.
	// Returns ErrUserNotFound if the user doesn't exist.
	GetByID(ctx context.Context, id int64) (*models.User, error)

	// GetByEmail retrieves a user by their email address.
	// Returns ErrUserNotFound if the user doesn't exist.
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

// ErrUserNotFound is returned when a user with the specified ID is not found.
var ErrUserNotFound = errors.New("user not found")

type usersRepository struct {
	dbpool *pgxpool.Pool
}

var _ UsersRepository = (*usersRepository)(nil)

func NewUsersRepository(dbpool *pgxpool.Pool) UsersRepository {
	return &usersRepository{
		dbpool: dbpool,
	}
}

func (r *usersRepository) Store(ctx context.Context, user *models.User) error {
	user.Email = r.normalizeEmail(user.Email)

	query := `
		INSERT INTO users (first_name, last_name, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at
	`
	err := r.dbpool.QueryRow(
		ctx,
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to store user: %w", err)
	}
	return nil
}

func (r *usersRepository) GetByID(ctx context.Context, id int64) (*models.User, error) {
	query := `
		SELECT id, first_name, last_name, email, password, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &models.User{}
	err := r.dbpool.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}

	return user, nil
}

func (r *usersRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	email = r.normalizeEmail(email)

	query := `
		SELECT id, first_name, last_name email, password, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &models.User{}
	err := r.dbpool.QueryRow(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}

	return user, nil
}

func (r *usersRepository) normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}
