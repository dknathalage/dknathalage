package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:  "dknathalage",
		Usage: "Utility cli for dknathalage",
		Commands: []*cli.Command{
			{
				Name:  "cloudrun",
				Usage: "cloud run commands",
				Subcommands: []*cli.Command{
					{
						Name:  "generate",
						Usage: "generate cloudrun configs",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "name",
								Aliases: []string{"n"},
								Usage:   "service name",
							},
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
