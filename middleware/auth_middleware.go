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
package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/alexisbcz/vauban/database/repositories"
	"github.com/alexisbcz/vauban/httperror"
	"github.com/alexisbcz/vauban/session"
)

// AuthMiddleware handles authentication for protected routes
type AuthMiddleware struct {
	usersRepository repositories.UsersRepository
}

// NewAuthMiddleware creates a new instance of the auth middleware
func NewAuthMiddleware(usersRepository repositories.UsersRepository) *AuthMiddleware {
	return &AuthMiddleware{
		usersRepository: usersRepository,
	}
}

// contextKey is a type for context keys to avoid collisions
type contextKey string

// UserKey is the key used to store the user in the request context
const UserKey contextKey = "user"

// Authenticate is a middleware function that checks if the user is authenticated
// and adds the user to the request context if they are
func (m *AuthMiddleware) Authenticate(next func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		// Get the authentication cookie
		cookieValue, ok := session.GetSignedCookie(r, session.AUTH_COOKIE_NAME)
		if !ok {
			slog.Debug("auth middleware: no valid auth cookie found")
			return httperror.Unauthorized("authentication required")
		}

		// Parse the user ID from the cookie
		userID, err := strconv.ParseInt(cookieValue, 10, 64)
		if err != nil {
			slog.Debug("auth middleware: invalid user ID in cookie", "err", err.Error())
			return httperror.Unauthorized("invalid authentication")
		}

		// Get the user from the database
		user, err := m.usersRepository.GetByID(r.Context(), userID)
		if err != nil {
			slog.Debug("auth middleware: user not found", "userID", userID, "err", err.Error())
			return httperror.Unauthorized("user not found")
		}

		// Add the user to the request context
		ctx := context.WithValue(r.Context(), UserKey, user)
		r = r.WithContext(ctx)

		// Call the next handler
		return next(w, r)
	}
}

// RequireAuth creates a protected route that requires authentication
func (m *AuthMiddleware) RequireAuth(handler func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) error {
	return m.Authenticate(handler)
}

// GetUserFromContext retrieves the authenticated user from the request context
// This can be used in your controllers to access the current user
func GetUserFromContext(r *http.Request) (interface{}, bool) {
	user := r.Context().Value(UserKey)
	if user == nil {
		return nil, false
	}
	return user, true
}

// Optional is a middleware that attempts to authenticate but doesn't require it
// This can be used for routes that work for both authenticated and unauthenticated users
func (m *AuthMiddleware) Optional(next func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		// Try to get the authentication cookie
		cookieValue, ok := session.GetSignedCookie(r, session.AUTH_COOKIE_NAME)
		if !ok {
			// Cookie exists, try to authenticate
			userID, err := strconv.ParseInt(cookieValue, 10, 64)
			if err == nil {
				user, err := m.usersRepository.GetByID(r.Context(), userID)
				if err == nil {
					// Add the user to the request context
					ctx := context.WithValue(r.Context(), UserKey, user)
					r = r.WithContext(ctx)
				}
			}
		}

		// Continue to the next handler regardless of authentication result
		return next(w, r)
	}
}
