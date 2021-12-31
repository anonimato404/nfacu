package cmd

import (
	"fmt"
	"github.com/UltiRequiem/nfacu/internal"
	"strings"
)

func Main() {

	_, configPath := getArguments()

	projectsConfig, errorGettingConfig := getConfig(configPath)

	if errorGettingConfig != nil {
		fmt.Printf("Error getting config: %s\n", errorGettingConfig.Error())
		return
	}

	fmt.Println(fmt.Sprintf(`Config read from "%s".`, configPath))

	for _, project := range projectsConfig {
		rawProjectConfig, errorGettingProjectConfig := getProjectConfig(project.Path)

		if errorGettingProjectConfig != nil {
			fmt.Printf("Error getting project config: %s\n", errorGettingProjectConfig.Error())
			return
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

		errorSavingAppConfig := saveConfigFile(project.Path, configRawData)

		if errorGettingProjectConfig != nil {
			fmt.Printf("Error saving project config: %s\n", errorSavingAppConfig.Error())
			return
		}

                fmt.Printf("%s updated successfully!\n", project.Path)

	}

}
