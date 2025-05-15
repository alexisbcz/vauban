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

package router

import (
	"errors"
	"fmt"
	"io/fs"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/alexisbcz/vauban/controllers"
	"github.com/alexisbcz/vauban/database/repositories"
	"github.com/alexisbcz/vauban/exit"
	"github.com/alexisbcz/vauban/httperror"
	"github.com/alexisbcz/vauban/middleware"
	"github.com/alexisbcz/vauban/public"
	"github.com/alexisbcz/vauban/session"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Middleware represents a chain of HTTP middleware
type Middleware func(handler HTTPHandlerWithErr) HTTPHandlerWithErr

// Chain combines multiple middleware into a single middleware
func Chain(middlewares ...Middleware) Middleware {
	return func(final HTTPHandlerWithErr) HTTPHandlerWithErr {
		for i := len(middlewares) - 1; i >= 0; i-- {
			final = middlewares[i](final)
		}
		return final
	}
}

type Router struct {
	*http.ServeMux
}

func New(dbpool *pgxpool.Pool) *Router {
	router := &Router{http.NewServeMux()}

	router.handlePublicAssets()
	router.handleHotReload()

	usersRepository := repositories.NewUsersRepository(dbpool)

	signUpController := controllers.NewSignUpController(usersRepository)
	router.Get("/sign-up/{$}", signUpController.Show)
	router.Post("/sign-up/{$}", signUpController.Handle)

	signInController := controllers.NewSignInController(usersRepository)
	router.Get("/sign-in/{$}", signInController.Show)
	router.Post("/sign-in/{$}", signInController.Handle)

	forgotPasswordController := controllers.NewForgotPasswordController(usersRepository)
	router.Get("/forgot-password/{$}", forgotPasswordController.Show)
	router.Post("/forgot-password/{$}", forgotPasswordController.Handle)

	resetPasswordController := controllers.NewResetPasswordController(usersRepository)
	router.Get("/reset-password/{$}", resetPasswordController.Show)
	router.Post("/reset-password/{$}", resetPasswordController.Handle)

	// Create middleware
	authMiddleware := middleware.NewAuthMiddleware(usersRepository)

	// Logout route
	router.Get("/logout/{$}", authMiddleware.RequireAuth(func(w http.ResponseWriter, r *http.Request) error {
		// Clear the auth cookie and redirect to home
		session.ClearCookie(w, session.AUTH_COOKIE_NAME)
		w.Header().Set("HX-Redirect", "/")
		return nil
	}))

	return router
}

func (r *Router) handlePublicAssets() {
	assetsFS, err := fs.Sub(public.FS, "assets")
	if err != nil {
		exit.WithErr(err)
	}
	r.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(assetsFS))))
}

var hotReloadOnlyOnce sync.Once

func (r *Router) handleHotReload() {
	r.Get("/hotreload", func(w http.ResponseWriter, req *http.Request) error {
		// Set headers for SSE
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Flush headers immediately
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		// Send ping events every 15 seconds to keep connection alive
		ticker := time.NewTicker(15 * time.Second)
		defer ticker.Stop()

		// Send initial ping to confirm connection is working
		fmt.Fprintf(w, "event: ping\ndata: connected\n\n")
		if f, ok := w.(http.Flusher); ok {
			f.Flush()
		}

		// Send reload event once if this is first connection after server start
		hotReloadOnlyOnce.Do(func() {
			time.Sleep(500 * time.Millisecond) // Wait a bit to ensure client is ready
			fmt.Fprintf(w, "event: reload\ndata: reload\n\n")
			if f, ok := w.(http.Flusher); ok {
				f.Flush()
			}
		})

		// Use a separate goroutine to send ping events
		done := make(chan bool)
		go func() {
			for {
				select {
				case <-ticker.C:
					fmt.Fprintf(w, "event: ping\ndata: ping\n\n")
					if f, ok := w.(http.Flusher); ok {
						f.Flush()
					}
				case <-done:
					return
				}
			}
		}()

		// Keep the connection open until client disconnects
		<-req.Context().Done()
		close(done)
		return nil
	})
}

type HTTPHandlerWithErr func(http.ResponseWriter, *http.Request) error

func (r *Router) handleFunc(pattern string, handler HTTPHandlerWithErr) {
	slog.Info("registering http route", "pattern", pattern)

	r.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			// Check if it's an HTTPError
			var httpErr *httperror.HTTPError
			if errors.As(err, &httpErr) {
				// Use the status code from the HTTPError
				http.Error(w, err.Error(), httpErr.Code)
				slog.Debug("http error", "code", httpErr.Code, "err", err.Error())
			} else {
				// Default to 500 for non-HTTP errors
				http.Error(w, err.Error(), http.StatusInternalServerError)
				slog.Error("internal server error", "err", err.Error())
			}
		}
	})
}

func (r *Router) Get(pattern string, handler HTTPHandlerWithErr) {
	r.handleFunc("GET "+pattern, handler)
}

func (r *Router) Post(pattern string, handler HTTPHandlerWithErr) {
	r.handleFunc("POST "+pattern, handler)
}

func (r *Router) Patch(pattern string, handler HTTPHandlerWithErr) {
	r.handleFunc("PATCH "+pattern, handler)
}

func (r *Router) Put(pattern string, handler HTTPHandlerWithErr) {
	r.handleFunc("PUT "+pattern, handler)
}

func (r *Router) Delete(pattern string, handler HTTPHandlerWithErr) {
	r.handleFunc("DELETE "+pattern, handler)
}
