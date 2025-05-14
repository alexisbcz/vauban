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
// Package session provides secure cookie-based session management functionality.
// It offers methods to create, validate, and clear signed HTTP cookies.
package session

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/alexisbcz/vauban/env"
)

// AUTH_COOKIE_NAME is the standard name used for authentication cookies
const AUTH_COOKIE_NAME = "vauban_auth_session"

// secretKey is used for HMAC signing of cookie values
// In production, this should be set via environment variable
var secretKey = []byte(env.GetVar("SESSION_KEY", "replace_this_key_in_prod"))

// SetSignedCookie creates and sets a secure, signed HTTP cookie
// The cookie value is signed with HMAC-SHA256 to prevent tampering
func SetSignedCookie(w http.ResponseWriter, name, value string, maxAge int) {
	// Create HMAC signature of the value
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(value))
	signature := base64.URLEncoding.EncodeToString(h.Sum(nil))

	// Combine signature and value for storage
	signedValue := signature + "|" + value

	// Set the cookie with secure defaults
	cookie := &http.Cookie{
		Name:     name,
		Value:    signedValue,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
}

// GetSignedCookie retrieves and validates a signed cookie
func GetSignedCookie(r *http.Request, name string) (string, bool) {
	// Attempt to get the cookie
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", false
	}

	// Split the cookie value to get signature and original value
	parts := strings.SplitN(cookie.Value, "|", 2)
	if len(parts) != 2 {
		return "", false
	}

	signature := parts[0]
	value := parts[1]

	// Verify the signature
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(value))
	expectedSignature := base64.URLEncoding.EncodeToString(h.Sum(nil))

	// Check if signatures match
	if signature != expectedSignature {
		return "", false
	}

	return value, true
}

// ClearCookie invalidates a cookie by setting it to expire immediately
func ClearCookie(w http.ResponseWriter, name string) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
	})
}
