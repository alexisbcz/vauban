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

package controllers

import (
	"net/http"

	"github.com/alexisbcz/vauban/httperror"
	"github.com/alexisbcz/vauban/views/pages/containers"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	docker "github.com/docker/docker/client"
)

type ContainersController struct{}

func NewContainersController() *ContainersController {
	return &ContainersController{}
}

func (c *ContainersController) Index(w http.ResponseWriter, r *http.Request) error {
	cli, err := docker.NewClientWithOpts(docker.FromEnv, docker.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer cli.Close()

	ctrs, err := cli.ContainerList(r.Context(), container.ListOptions{All: true})
	if err != nil {
		return err
	}

	services, err := cli.ServiceList(r.Context(), types.ServiceListOptions{})
	if err != nil {
		return err
	}

	return containers.IndexPage(ctrs, services).Render(w)
}

func (c *ContainersController) Show(w http.ResponseWriter, r *http.Request) error {
	return httperror.NotFound("container not found")
}
