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
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/alexisbcz/vauban/env"
	"github.com/alexisbcz/vauban/exit"
	"github.com/alexisbcz/vauban/router"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbpool, err := pgxpool.New(context.Background(), env.GetVar("DATABASE_URL", "postgres://vauban:vauban@localhost/vauban?sslmode=disable"))
	if err != nil {
		exit.WithErr(err)
	}
	defer dbpool.Close()

	if err := dbpool.Ping(context.Background()); err != nil {
		exit.WithErr(err)
	}

	r := router.New(dbpool)
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
