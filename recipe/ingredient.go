package recipe

import (
	"fmt"

	"github.com/goccy/go-yaml"
)

type IngredientType string

const (
	IngredientTemplate IngredientType = "template"
	IngredientSymlink  IngredientType = "symlink"
	IngredientCopy     IngredientType = "copy"
	IngredientRun      IngredientType = "run"
)

func (t IngredientType) Valid() bool {
	switch t {
	case IngredientTemplate,
		IngredientSymlink,
		IngredientCopy,
		IngredientRun:
		return true
	default:
		return false
	}
}

func (t *IngredientType) UnmarshalYAML(data []byte) error {
	var value string
	if err := yaml.Unmarshal(data, &value); err != nil {
		return err
	}

	parsed := IngredientType(value)
	if !parsed.Valid() {
		return fmt.Errorf(
			"invalid ingredient of type %q; expected one of %q, %q, %q, %q",
			value,
			IngredientCopy,
			IngredientRun,
			IngredientSymlink,
			IngredientTemplate,
		)
	}

	*t = parsed
	return nil
}

type Ingredient struct {
	Type IngredientType `yaml:"type"`

	Source  string   `yaml:"source,omitempty"`
	Target  string   `yaml:"target,omitempty"`
	Command string   `yaml:"command,omitempty"`
	Args    []string `yaml:"args, omitempty"`
}
