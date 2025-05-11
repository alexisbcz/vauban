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
	"fmt"
	"log/slog"
	"net/http"

	"github.com/alexisbcz/vauban/pkg/controllers"
)

type Router struct {
	*http.ServeMux
}

func New() *Router {
	router := &Router{http.NewServeMux()}

	containersController := controllers.NewContainersController()
	router.Get("/containers/{$}", containersController.Index)

	return router
}

type HTTPHandlerWithErr func(http.ResponseWriter, *http.Request) error

func (r *Router) Handle(pattern string, handler HTTPHandlerWithErr) {
	slog.Info("registering http route", "pattern", pattern)
	r.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		if err := handler(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

func (r *Router) Get(pattern string, handler HTTPHandlerWithErr) {
	r.Handle("GET "+pattern, handler)
}

func (r *Router) Post(pattern string, handler HTTPHandlerWithErr) {
	r.Handle("POST "+pattern, handler)
}

func (r *Router) Patch(pattern string, handler HTTPHandlerWithErr) {
	r.Handle("PATCH "+pattern, handler)
}

func (r *Router) Put(pattern string, handler HTTPHandlerWithErr) {
	r.Handle("PUT "+pattern, handler)
}

func (r *Router) Delete(pattern string, handler HTTPHandlerWithErr) {
	r.Handle("DELETE "+pattern, handler)
}
