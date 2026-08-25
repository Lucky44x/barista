package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:                  "barista",
		EnableShellCompletion: true,
		Suggest:               true,
		Commands: []*cli.Command{
			{
				Name:                  "brew",
				Usage:                 "barista brew <recipe-yaml> [options]",
				Description:           "Runs the defined recipe",
				EnableShellCompletion: true,
				Suggest:               true,
				Arguments: []cli.Argument{
					&cli.StringArg{
						Name:      "Recipe-file",
						UsageText: "The Recipe `FILE` that should be executed",
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
				Action: func(context.Context, *cli.Command) error {
					fmt.Println("Hello World")
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
