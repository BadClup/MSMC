package shared

import (
	"context"
	"encoding/json"
	"os"
)

type MinecraftEngine string

func (e MinecraftEngine) String() string {
	return string(e)
}

// enum implementation:
const (
	VanillaEngine MinecraftEngine = "vanilla"
	ForgeEngine   MinecraftEngine = "forge"
)

// ServerInstance default tag is used by swagger
type ServerInstance struct {
	DockerContainerID string          `json:"docker_id" default:"af5bb532db04"`
	Name              string          `json:"name" default:"My Server"`
	Port              int             `json:"port" default:"25565"`
	Engine            MinecraftEngine `json:"engine" default:"vanilla"`
	Version           string          `json:"version" default:"1.16.5"`
	Seed              string          `json:"seed" default:"example-seed"`
}

func ReadServerInstances() ([]ServerInstance, error) {
	file, err := os.ReadFile(InstancesPath)
	if err != nil {
		return nil, err
	}

	var servers []ServerInstance
	return servers, json.Unmarshal(file, &servers)
}

func (i ServerInstance) SaveInInstancesList() error {
	instances, err := ReadServerInstances()
	if err != nil {
		return err
	}

	instances = append(instances, i)

	file, err := json.Marshal(instances)
	if err != nil {
		return err
	}

	return os.WriteFile(InstancesPath, file, 0644)
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
	Running        bool `json:"running" default:"false"`
}
