package engine

import (
	"os"
	"strings"

	"github.com/flosch/pongo2/v7"
)

func CollectEnvVars() pongo2.Context {
	allEnvVars := os.Environ()

	envMap := make(map[string]string)
	for _, ev := range allEnvVars {
		pair := strings.SplitN(ev, "=", 2)
		if len(pair) == 2 {
			envMap[pair[0]] = pair[1]
		}
	}

	return pongo2.Context{
		"env": envMap,
	}
}
