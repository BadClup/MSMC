package server

import (
	"backend/shared"
	"encoding/json"
	"fmt"
	"os"
)

func getAllServers() ([]shared.ServerInstanceStatus, error) {
	file, err := os.ReadFile(shared.InstancesPath)
	if err != nil {
		return nil, err
	}

	var servers []shared.ServerInstance
	if err := json.Unmarshal(file, &servers); err != nil {
		return nil, err
	}

	err = ensureServersContainersExist(servers)
	if err != nil {
		return nil, err
	}

	serversStatus := make([]shared.ServerInstanceStatus, len(servers))
	for _, server := range servers {
		status, err := server.GetStatus()
		if err != nil {
			return nil, err
		}
		serversStatus = append(serversStatus, status)
	}

	return serversStatus, nil
}

func ensureServersContainersExist(servers []shared.ServerInstance) error {
	var containersIds []string
	for _, server := range servers {
		containersIds = append(containersIds, server.DockerContainerID)
	}

	existingContainers, err := shared.GetContainers(containersIds)
	if err != nil {
		return err
	}

	if len(existingContainers) != len(servers) {
		// TODO: handle this case
		_, _ = fmt.Fprintf(os.Stderr, "WARING: %d servers were probably deleted by user or external program, we cannot do anything about it yet\n", len(servers)-len(existingContainers))
	}

	return nil
}
