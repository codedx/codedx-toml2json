package main

import (
	"flag"
	"fmt"
	"github.com/codedx/codedx-toml2json/pkg/console"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	/* flag.go uses exit code 2 for invalid command-line arguments */
	invalidCommandLineArgumentsExitCode = 3
	expectedTOMLFileExtensionExitCode   = 4
	jsonConfigFileAlreadyExistsExitCode = 5
	cannotReadTOMLConfigExitCode        = 6
	cannotWriteJSONConfigExitCode       = 7
)

func main() {

	const tomlFilePathFlagName = "tomlFile"
	const jsonFilePathFlagName = "jsonFile"

	tomlFilePathFlagValue := flag.String(tomlFilePathFlagName, "", "a path to the TOML file to convert to JSON")
	jsonFilePathFlagValue := flag.String(jsonFilePathFlagName, "", "a path to the JSON ouptut file")

	flag.Parse()

	tomlFilePath := console.ReadFileFlagValue(tomlFilePathFlagName, tomlFilePathFlagValue, true, invalidCommandLineArgumentsExitCode)
	jsonFilePath := console.ReadRequiredFlagStringValue(jsonFilePathFlagName, jsonFilePathFlagValue, invalidCommandLineArgumentsExitCode)

	jsonExtension := filepath.Ext(jsonFilePath)
	if jsonExtension != ".json" {
		console.Fatalf(invalidCommandLineArgumentsExitCode, "JSON file must have '.json' extension. Extension '%s' is unsupported.", jsonExtension)
	}

	tomlDirectory, tomlFilename := filepath.Split(tomlFilePath)
	tomlFilenameExt := filepath.Ext(tomlFilename)

	if tomlFilenameExt != ".toml" {
		console.Fatalf(expectedTOMLFileExtensionExitCode, "Expected TOML file to have '.toml' file extension, found %s", tomlFilenameExt)
	}
	tomlFilenameNoExt := tomlFilename[0 : len(tomlFilename)-len(tomlFilenameExt)]

	jsonSiblingPath := filepath.Join(tomlDirectory, fmt.Sprintf("%s.json", tomlFilenameNoExt))
	if _, err := os.Stat(jsonSiblingPath); !os.IsNotExist(err) {
		console.Fatalf(jsonConfigFileAlreadyExistsExitCode, "TOML to JSON conversion will not work because %s will block reading %s. Retry after either deleting the JSON file or specifying another path.", jsonSiblingPath, tomlFilePath)
	}

	viper.SetConfigName(tomlFilenameNoExt)
	viper.SetConfigType("toml")
	viper.AddConfigPath(tomlDirectory)

	if err := viper.ReadInConfig(); err != nil {
		console.Fatalf(cannotReadTOMLConfigExitCode, "Cannot read TOML configuration from %s: %s", tomlFilePath, err.Error())
	}

	if err := viper.WriteConfigAs(jsonFilePath); err != nil {
		console.Fatalf(cannotWriteJSONConfigExitCode, "Cannot write JSON configuration to '%s': %s", jsonFilePath, err.Error())
	}
}
