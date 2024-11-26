package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v66/github"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GithubToken string `envconfig:"GITHUB_TOKEN"`
}

func LoadConfig() Config {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func main() {
	cfg := LoadConfig()
	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(string(cfg.GithubToken))

	user, resp, err := client.Users.Get(ctx, "")
	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return
	}

	// Rate.Limit should most likely be 5000 when authorized.
	log.Printf("Rate: %#v\n", resp.Rate)

	// If a Token Expiration has been set, it will be displayed.
	if !resp.TokenExpiration.IsZero() {
		log.Printf("Token Expiration: %v\n", resp.TokenExpiration)
	}

	fmt.Printf("\n%v\n", github.Stringify(user))
}
