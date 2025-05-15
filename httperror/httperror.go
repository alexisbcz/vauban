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

package httperror

import (
	"errors"
	"net/http"
)

// HTTPError represents an error that can occur while handling a request.
type HTTPError struct {
	error
	Code int
}

// New creates an instance of HTTPError
func New(code int, message string) *HTTPError {
	return &HTTPError{
		error: errors.New(message),
		Code:  code,
	}
}

// BadRequest creates an instance of HTTPError, setting code to 400
func BadRequest(message string) *HTTPError {
	return New(http.StatusBadRequest, message)
}

// InternalServerError creates an instance of HTTPError, setting code to 500
func InternalServerError(message string) *HTTPError {
	return New(http.StatusInternalServerError, message)
}

// NotFound creates an instance of HTTPError, setting code to 404
func NotFound(message string) *HTTPError {
	return New(http.StatusNotFound, message)
}

// NotImplemented creates an instance of HTTPError, setting code to 501
func NotImplemented(message string) *HTTPError {
	return New(http.StatusNotImplemented, message)
}

// Unauthorized creates an instance of HTTPError, setting code to 401
func Unauthorized(message string) *HTTPError {
	return New(http.StatusUnauthorized, message)
}
