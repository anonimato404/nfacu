package cmd

import (
	"fmt"
	"github.com/UltiRequiem/nfacu/internal"
	"strings"
)

func Main() {
	projectsConfig, errorGettingConfig := getConfig("demo.json")

	if errorGettingConfig != nil {
		fmt.Printf("Error getting config: %s\n", errorGettingConfig.Error())
	}

	for _, project := range projectsConfig {
		rawProjectConfig, errorGettingProjectConfig := getProjectConfig(project.Path)

		if errorGettingProjectConfig != nil {
			fmt.Printf("Error getting project config: %s\n", errorGettingProjectConfig.Error())
		}

		configRawData := ""

		for _, line := range rawProjectConfig {
			for key := range project.Settings {
				if strings.Contains(line, fmt.Sprintf(`"%s"`, key)) {
					configRawData += internal.ParseLine(key, project.Settings[key])
				} else {
					configRawData += line
				}

				configRawData += "\n"
				break
			}
		}

	}

}
