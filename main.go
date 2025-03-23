package main

import (
	"fmt"

	"github.com/dknathalage/dknathalage/internal"
	_ "github.com/dknathalage/dknathalage/plugins" // Register plugins
)

func main() {
	config, err := internal.LoadConfig()
	if err != nil {
		fmt.Println(err)
		return
	}

	mkdown, err := internal.GenerateMarkdown()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = internal.SaveMarkdown(config, mkdown)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = internal.CheckFileSize(config)
	if err != nil {
		fmt.Println(err)
		return
	}
}
