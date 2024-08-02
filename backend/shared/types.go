package shared

import (
	"context"
)

type MinecraftEngine string

// enum implementation:
const (
	VanillaEngine MinecraftEngine = "vanilla"
	ForgeEngine   MinecraftEngine = "forge"
)

type ServerInstance struct {
	DockerContainerID string          `json:"docker_id"`
	Name              string          `json:"name"`
	Port              string          `json:"port"`
	Engine            MinecraftEngine `json:"engine"`
	Seed              string          `json:"seed"`
}

func (i ServerInstance) GetStatus() (ServerInstanceStatus, error) {
	container, err := DockerClient.ContainerInspect(context.Background(), i.DockerContainerID)
	if err != nil {
		return ServerInstanceStatus{}, err
	}

	return ServerInstanceStatus{
		ServerInstance: i,
		Running:        container.State.Running,
	}, nil
}

type ServerInstanceStatus struct {
	ServerInstance `json:"server_instance"`
	Running        bool `json:"running"`
}
