package shared

import (
	"context"
	"fmt"
	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"os"
)

const MinecraftImageName = "itzg/minecraft-server"

var DockerClient = func() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		fmt.Println("Failed to create docker client")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cli
}()

func GetContainers(containersIds []string) ([]dockerTypes.ContainerJSON, error) {
	ctx := context.Background()

	var containers []dockerTypes.ContainerJSON
	for _, id := range containersIds {
		container, err := DockerClient.ContainerInspect(ctx, id)
		if err != nil {
			return nil, err
		}
		containers = append(containers, container)
	}

	return containers, nil
}

func EnsureDockerProperlyInitialized() error {
	// TODO: check if docker installed
	// TODO: check if required directories are installed
	return pullNeededImages()
}

func pullNeededImages() error {
	ctx := context.Background()

	imagesToPull := []string{
		MinecraftImageName,
	}

	for _, imageName := range imagesToPull {
		reader, err := DockerClient.ImagePull(ctx, imageName, image.PullOptions{})
		if err != nil {
			return err
		}

		if err := reader.Close(); err != nil {
			return err
		}
	}

	return nil
}
