package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name: "dknathalage",
		Commands: []*cli.Command{
			{
				Name: "cloudrun",
				Subcommands: []*cli.Command{
					{
						Name: "generate",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "name", Aliases: []string{"n"}},
						},
					},
				},
			},
		},
	}
}

func main() {
	NewApp().Run(os.Args)
}
