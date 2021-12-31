package cmd

import (
	"fmt"
	"github.com/UltiRequiem/nfacu/internal"
	"os"
	"strings"
)

func Main() {
	projectsConfig, errorGettingConfig := getConfig("demo.json")

	if errorGettingConfig != nil {
		fmt.Printf("Error getting config: %s\n", errorGettingConfig.Error())
	}

	for _, project := range projectsConfig {

		projectConfig, errorReadingProjectConfig := os.ReadFile(project.Path)

		if errorReadingProjectConfig != nil {
			fmt.Printf("Error reading project %s config: %s\n", project.Path, errorReadingProjectConfig)
		}

		temp := strings.Split(string(projectConfig), "\n")

		configRawData := ""

		for _, line := range temp {
			for key := range project.Settings {
				if strings.Contains(line, fmt.Sprintf(`"%s"`, key)) {
					configRawData += internal.ParseLine(key, project.Settings[key]) + "\n"
					break
				} else {
					configRawData += line + "\n"
					break
				}
			}
		}

		errorWritingProjectConfig := os.WriteFile(project.Path, []byte(configRawData), 0644)

		if errorWritingProjectConfig != nil {
			fmt.Printf("Error writing project %s config: %s\n", project.Path, errorWritingProjectConfig)
		}
	}

}
