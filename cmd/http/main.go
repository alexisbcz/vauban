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

package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/alexisbcz/vauban/pkg/env"
	"github.com/alexisbcz/vauban/pkg/router"
)

func main() {
	r := router.New()
	srv := &http.Server{
		Addr:    env.GetVar("HTTP_ADDR", ":8080"),
		Handler: r,
	}

	slog.Info("starting http server", "addr", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		slog.Error("some error occured while serving http", "err", err)
		os.Exit(1)
	}
}
