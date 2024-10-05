package main

import (
	"log"
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
				Usage: "add a task to the list",
			},
		},
	}
}

func main() {
	app := NewApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
