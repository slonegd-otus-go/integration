package docker

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

// отсюда https://medium.com/backendarmy/controlling-the-docker-engine-in-go-d25fc0fe2c45
func StartNewContainer(cli *client.Client, image string) (string, error) {
	hostBinding := nat.PortBinding{
		HostIP:   "0.0.0.0",
		HostPort: "8080",
	}
	containerPort, err := nat.NewPort("tcp", "80")
	if err != nil {
		return "", fmt.Errorf("new port failed: %s", err)
	}

	portBinding := nat.PortMap{containerPort: []nat.PortBinding{hostBinding}}
	cont, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: image,
		},
		&container.HostConfig{
			PortBindings: portBinding,
		}, nil, "")
	if err != nil {
		return "", fmt.Errorf("create container failed: %s", err)
	}

	err = cli.ContainerStart(context.Background(), cont.ID, types.ContainerStartOptions{})
	if err != nil {
		return "", fmt.Errorf("start container failed: %s", err)
	}

	return cont.ID, nil
}

func StopContainer(cli *client.Client, id string) error {
	timeout := time.Duration(1 * time.Second)
	err := cli.ContainerStop(context.Background(), id, &timeout)
	if err != nil {
		return fmt.Errorf("stop container failed: %s", err)
	}
	return nil
}
