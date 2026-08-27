package commands

import (
	"context"
	"fmt"

	"github.com/lucky44x/barista/recipe"
	"github.com/urfave/cli/v3"
)

func Taste(ctx context.Context, cmd *cli.Command) error {
	recipeFile := cmd.StringArg("Recipe-file")
	recipe, err := recipe.ReadRecipe(recipeFile)
	if err != nil {
		return err
	}

	fmt.Println(recipe)

	return nil
}
