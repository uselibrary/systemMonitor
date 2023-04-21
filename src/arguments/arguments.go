package arguments

import (
	"os"
)

// function to get command line arguments
// get command line arguments
func GetCommandArguments() string {
	// get command line arguments
	arguments := os.Args
	// get config file path
	configFilePath := ""
	// get command line arguments length
	argumentsLength := len(arguments)
	// get command line arguments
	for i := 1; i < argumentsLength; i++ {
		// get config file path
		if arguments[i] == "--config" || arguments[i] == "-c" {
			// get config file path
			configFilePath = arguments[i+1]
			// break
			break
		}
	}

	return configFilePath
}
