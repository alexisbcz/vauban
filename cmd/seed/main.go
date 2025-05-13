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
	"time"

	"github.com/alexisbcz/vauban/database/seeders"
	"github.com/alexisbcz/vauban/env"
	"github.com/alexisbcz/vauban/exit"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, env.GetVar("DATABASE_URL", "postgres://panache:panache@localhost/panache?sslmode=disable"))
	if err != nil {
		exit.WithErr(err)
	}
	defer dbpool.Close()

	seeders := []seeders.Seeder{
		seeders.NewUserSeeder(dbpool),
	}

	for _, seeder := range seeders {
		if err := seeder.Seed(ctx); err != nil {
			exit.WithErr(err)
		}
	}
}
