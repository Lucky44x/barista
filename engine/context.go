package engine

import (
	"os"
	"strings"
)

func CollectEnvVars() map[string]string {
	allEnvVars := os.Environ()

	envMap := make(map[string]string)
	for _, ev := range allEnvVars {
		pair := strings.SplitN(ev, "=", 2)
		if len(pair) == 2 {
			envMap[pair[0]] = pair[1]
		}
	}

	return envMap
}

func CollectInputs() map[string]any {
	return nil
}
