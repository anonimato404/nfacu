package cmd

import (
	"encoding/json"
	"github.com/UltiRequiem/nfacu/internal"
	"os"
)

func getConfig(configFilePath string) (internal.NFACUConfig, error) {
	configData, errorReadingConfig := os.ReadFile(configFilePath)

	if errorReadingConfig != nil {
		return nil, errorReadingConfig
	}

	var projectsConfig internal.NFACUConfig

	errorUnmarshalling := json.Unmarshal(configData, &projectsConfig)

	if errorUnmarshalling != nil {
		return nil, errorReadingConfig
	}

	return projectsConfig, nil
}
