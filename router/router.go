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
	"io/fs"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/alexisbcz/vauban/controllers"
	"github.com/alexisbcz/vauban/database/repositories"
	"github.com/alexisbcz/vauban/exit"
	"github.com/alexisbcz/vauban/httperror"
	"github.com/alexisbcz/vauban/public"
	"github.com/jackc/pgx/v5/pgxpool"
	datastar "github.com/starfederation/datastar/sdk/go"
)

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

	containersController := controllers.NewContainersController()
	router.Get("/containers/{$}", containersController.Index)
	router.Get("/containers/{id}/{$}", containersController.Show)

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
	r.HandleFunc("/hotreload", func(w http.ResponseWriter, req *http.Request) {
		sse := datastar.NewSSE(w, req)
		hotReloadOnlyOnce.Do(func() {
			// Refresh the client page as soon as connection
			// is established. This will occur only once
			// after the server starts.
			sse.ExecuteScript(
				"window.location.reload()",
				datastar.WithExecuteScriptRetryDuration(time.Second),
			)
		})

		// Freeze the event stream until the connection
		// is lost for any reason. This will force the client
		// to attempt to reconnect after the server reboots.
		<-req.Context().Done()
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
