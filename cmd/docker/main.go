package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// List all containers
	containers, err := cli.ContainerList(ctx, container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	// Group containers by compose service
	composeServices := groupContainersByComposeService(containers)

	// Print the containers grouped by service
	fmt.Println("Containers")
	fmt.Println("=====================================")

	// Sort service names for consistent output
	var serviceNames []string
	for service := range composeServices {
		serviceNames = append(serviceNames, service)
	}
	sort.Strings(serviceNames)

	for _, service := range serviceNames {
		containers := composeServices[service]
		fmt.Printf("\nService: %s (%d containers)\n", service, len(containers))
		fmt.Println("-----------------------------------")

		// Get the project filepath and config files if they exist (should be same for all containers in a service)
		if service != "non-compose" && len(containers) > 0 {
			// Get project working directory
			if path, exists := containers[0].Labels["com.docker.compose.project.working_dir"]; exists {
				fmt.Printf(" Project path: %s\n", path)
			}
			// Get compose config files
			if configFiles, exists := containers[0].Labels["com.docker.compose.project.config_files"]; exists {
				fmt.Printf(" Config files: %s\n", configFiles)
			}
		}

		for _, c := range containers {
			fmt.Printf(" - %s (ID: %.12s)\n", c.Names[0], c.ID)
			fmt.Printf("   Status: %s\n", c.Status)
			fmt.Printf("   Image: %s\n", c.Image)
		}
	}
	fmt.Println()
	fmt.Println()

	// Pull Alpine image
	fmt.Println("Pulling Alpine image...")
	reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", image.PullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	// Create container with Alpine image - just create with a sleep command
	// so it stays alive for us to execute commands in it
	containerName := "alpine-demo"
	fmt.Println("\nCreating container:", containerName)
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"sleep", "60"}, // Container will stay alive for 60 seconds
		Tty:   true,
	}, nil, nil, nil, containerName)
	if err != nil {
		panic(err)
	}
	containerID := resp.ID
	fmt.Printf("Container created with ID: %.12s\n", containerID)

	// Start container
	fmt.Println("Starting container...")
	if err := cli.ContainerStart(ctx, containerID, container.StartOptions{}); err != nil {
		panic(err)
	}
	fmt.Println("Container started successfully")

	// Now let's execute various commands in the running container
	fmt.Println("\nExecuting commands in the running container:")

	// Execute ls -la / command
	fmt.Println("\n1. Directory listing (ls -la /)")
	fmt.Println("=====================================")
	execConfig := container.ExecOptions{
		Cmd:          []string{"ls", "-la", "/"},
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
	}
	execID, err := cli.ContainerExecCreate(ctx, containerID, execConfig)
	if err != nil {
		panic(err)
	}

	// Attach to the exec instance to get its output
	execAttachResp, err := cli.ContainerExecAttach(ctx, execID.ID, container.ExecAttachOptions{})
	if err != nil {
		panic(err)
	}
	defer execAttachResp.Close()

	// Copy the output to stdout
	stdcopy.StdCopy(os.Stdout, os.Stderr, execAttachResp.Reader)

	// Execute command to show system info
	fmt.Println("\n2. System Information")
	fmt.Println("=====================================")
	execConfig = container.ExecOptions{
		Cmd:          []string{"sh", "-c", "cat /etc/os-release && echo '' && uname -a"},
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
	}
	execID, err = cli.ContainerExecCreate(ctx, containerID, execConfig)
	if err != nil {
		panic(err)
	}

	// Attach to the exec instance
	execAttachResp, err = cli.ContainerExecAttach(ctx, execID.ID, container.ExecAttachOptions{})
	if err != nil {
		panic(err)
	}
	defer execAttachResp.Close()

	// Copy the output to stdout
	stdcopy.StdCopy(os.Stdout, os.Stderr, execAttachResp.Reader)

	// Execute a command to create a file and then read it
	fmt.Println("\n3. Creating and reading a file")
	fmt.Println("=====================================")
	execConfig = container.ExecOptions{
		Cmd:          []string{"sh", "-c", "echo 'Hello from Alpine container!' > /tmp/hello.txt && cat /tmp/hello.txt"},
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
	}
	execID, err = cli.ContainerExecCreate(ctx, containerID, execConfig)
	if err != nil {
		panic(err)
	}

	// Attach to the exec instance
	execAttachResp, err = cli.ContainerExecAttach(ctx, execID.ID, container.ExecAttachOptions{})
	if err != nil {
		panic(err)
	}
	defer execAttachResp.Close()

	// Copy the output to stdout
	stdcopy.StdCopy(os.Stdout, os.Stderr, execAttachResp.Reader)

	// Stop and remove the container when we're done
	fmt.Println("\nStopping and removing container...")
	err = cli.ContainerStop(ctx, containerID, container.StopOptions{})
	if err != nil {
		panic(err)
	}

	err = cli.ContainerRemove(ctx, containerID, container.RemoveOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Container stopped and removed successfully")
}

// groupContainersByComposeService groups containers by their Docker Compose service name
func groupContainersByComposeService(containers []container.Summary) map[string][]container.Summary {
	serviceMap := make(map[string][]container.Summary)
	for _, c := range containers {
		// Look for the Docker Compose service label
		serviceName, exists := c.Labels["com.docker.compose.service"]
		if !exists {
			// If the container isn't part of a compose service, group it under "non-compose"
			serviceName = "non-compose"
		}
		serviceMap[serviceName] = append(serviceMap[serviceName], c)
	}
	return serviceMap
}
