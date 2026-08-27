package recipe

const TemplateIngredient = "template"
const SymlinkIngredient = "symlink"
const CopyIngredient = "copy"
const CommandIngredient = "run"

type Ingredient struct {
	Type string `yaml:"type"`

	Source  string   `yaml:"source,omitempty"`
	Target  string   `yaml:"target,omitempty"`
	Command string   `yaml:"command,omitempty"`
	Args    []string `yaml:"args, omitempty"`
}
