package recipe

import (
	"os"

	"github.com/goccy/go-yaml"
)

type Meta struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}

type Recipe struct {
	Meta  Meta   `yaml:"meta"`
	Steps []Step `yaml:"steps"`
}

func ReadRecipe(recipeFile string) (*Recipe, error) {
	var recipe Recipe
	bytes, ferr := os.ReadFile(recipeFile)
	if ferr != nil {
		return nil, ferr
	}

	berr := yaml.Unmarshal(bytes, &recipe)
	if berr != nil {
		return nil, berr
	}

	return &recipe, nil
}
