package recipe

type Step struct {
	Name        string       `yaml:"name"`
	Description string       `yaml:"description,omitempty"`
	Ingredients []Ingredient `yaml:"ingredients"`
}
