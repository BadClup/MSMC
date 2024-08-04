package server

import (
	"backend/shared"
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"math/rand"
	"strconv"
	"strings"
)

// default tags are being used by swagger docs
type createServerDto struct {
	Name    string                  `json:"name" default:"example-minecraft-server"`
	Port    *int                    `json:"port,omitempty" default:"25565"`
	Seed    *string                 `json:"seed,omitempty" default:"example-seed"`
	Version *string                 `json:"version,omitempty" default:"1.16.5"`
	Engine  *shared.MinecraftEngine `json:"engine,omitempty" default:"vanilla"`
}

func (d *createServerDto) intoServerInstance(dockerContainerId string) shared.ServerInstance {
	d.initializeEmpty()

	return shared.ServerInstance{
		DockerContainerID: dockerContainerId,
		Name:              d.Name,
		Port:              *d.Port,
		Engine:            *d.Engine,
		Version:           *d.Version,
		Seed:              *d.Seed,
	}
}

func (d *createServerDto) initializeEmpty() {
	if d.Port == nil {
		p := 25565
		d.Port = &p
	}

	if d.Engine == nil {
		e := shared.VanillaEngine
		d.Engine = &e
	}

	if d.Version == nil {
		v := shared.DefaultVersion
		d.Version = &v
	}

	if d.Seed == nil {
		s := ""
		d.Seed = &s
	}
}

func createServer(dto createServerDto) shared.ApiError {
	dto.initializeEmpty()

	instances, err := shared.ReadServerInstances()
	if err != nil {
		return shared.ApiErrorInternal(err)
	}

	for _, instance := range instances {
		if instance.Name == dto.Name {
			errMsg := fmt.Sprintf("server with name %s already exists", dto.Name)
			return shared.ApiErrorFromString(errMsg, 400)
		}
	}

	containerConfig := &container.Config{
		Image: shared.DockerMinecraftImage,
		ExposedPorts: nat.PortSet{
			"25565/tcp": struct{}{},
		},
		Env: []string{
			"EULA=true",
			fmt.Sprintf("VERSION=%s", *dto.Version),
			fmt.Sprintf("SEED=%s", *dto.Seed),
			fmt.Sprintf("TYPE=%s", strings.ToUpper(dto.Engine.String())),
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			nat.Port("25565/tcp"): []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: strconv.Itoa(*dto.Port),
				},
			},
		},
	}

	containerId := fmt.Sprintf("MSMC_server-%s-%s", dto.Name, randStringBytes(5))

	mcContainer, err := shared.DockerClient.ContainerCreate(
		context.Background(),
		containerConfig,
		hostConfig,
		nil,
		nil,
		containerId,
	)
	if err != nil {
		return shared.ApiErrorInternal(err)
	}

	serviceInstance := dto.intoServerInstance(mcContainer.ID)

	err = serviceInstance.SaveInInstancesList()
	return shared.ApiErrorInternal(err)
}

func randStringBytes(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
