package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func Main() {
	configData, errorReadingConfig := os.ReadFile("demo.json")

	if errorReadingConfig != nil {
		fmt.Printf("Error reading NFACU config file: %s\n", errorReadingConfig)
	}

	var projectsConfig NFACUConfig

	errorUnmarshalling := json.Unmarshal(configData, &projectsConfig)

	if errorUnmarshalling != nil {
		fmt.Printf("Error parsing NFACU config: %s\n", errorUnmarshalling)
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
					configRawData += parseLine(key, project.Settings[key]) + "\n"
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

func parseLine(key, val string) string {
	return fmt.Sprintf(` <add key="%s" value="%s" />`, key, val)

}
