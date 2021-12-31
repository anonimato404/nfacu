package cmd

import (
	"fmt"
	"strings"

	"github.com/UltiRequiem/nfacu/internal"
)

func Main() {

	_, configPath := getArguments()

	projectsConfig, errorGettingConfig := getConfig(configPath)

	if errorGettingConfig != nil {
		fmt.Printf("Error getting config: %s\n", errorGettingConfig.Error())
		return
	}

	fmt.Printf(`Config read from "%s".`+"\n", configPath)

	for _, project := range projectsConfig {
		rawProjectConfig, errorGettingProjectConfig := getProjectConfig(project.Path)

		if errorGettingProjectConfig != nil {
			fmt.Printf("Error getting project config: %s\n", errorGettingProjectConfig.Error())
			return
		}

		fmt.Printf("\n"+`Reading %s `+"\n", project.Path)

		configRawData := ""

		for _, line := range rawProjectConfig {
			lineAdded := false
			for key := range project.Settings {
				if strings.Contains(line, fmt.Sprintf(`"%s"`, key)) {
					configRawData += internal.ParseLine(key, project.Settings[key])
					lineAdded = true
					fmt.Printf(`Changing "%s" property to "%s" on "%s".`+"\n", key, project.Settings[key], project.Path)
				}
			}

			if !lineAdded {
				configRawData += line
			}

			configRawData += "\n"
		}

		errorSavingAppConfig := saveConfigFile(project.Path, configRawData)

		if errorGettingProjectConfig != nil {
			fmt.Printf("Error saving project config: %s\n", errorSavingAppConfig.Error())
			return
		}

		fmt.Printf("\n%s updated successfully!\n", project.Path)

	}

}
