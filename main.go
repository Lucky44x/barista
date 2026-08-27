package main

import (
	"context"
	"log"
	"os"

	"github.com/lucky44x/barista/commands"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:                  "barista",
		EnableShellCompletion: true,
		Suggest:               true,
		Commands: []*cli.Command{
			{
				Name:                  "taste",
				Usage:                 "barista taste",
				Description:           "Validates a given configuration",
				EnableShellCompletion: true,
				Suggest:               true,
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name:      "Recipe-file",
						UsageText: "<recipe-file>",
					},
				},
				Action: commands.Taste,
			},
			{
				Name:                  "brew",
				Usage:                 "barista brew",
				Description:           "Runs the defined recipe",
				EnableShellCompletion: true,
				Suggest:               true,
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name:      "Recipe-file",
						UsageText: "<recipe-file>",
					},
				},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "noParallel",
						Aliases: []string{"np"},
						Usage:   "Disable parallel execution of the recipe",
					},
					&cli.StringSliceFlag{
						Name:    "input",
						Aliases: []string{"i"},
						Usage:   "One (or multiple) inputs that should be supplied to the recipe",
					},
					&cli.StringSliceFlag{
						Name:    "inputFile",
						Aliases: []string{"if"},
						Usage:   "A yaml that defines inputs that should be supplied to the recipe",
					},
				},
				Action: commands.Brew,
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
