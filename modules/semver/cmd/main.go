package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/go-github/v66/github"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GithubToken      string `envconfig:"GITHUB_TOKEN" required:"true"`
	GithubRepoOwner  string `envconfig:"GITHUB_REPO_OWNER" required:"true"`
	GithubRepository string `envconfig:"GITHUB_REPOSITORY" required:"true"`
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

	user, resp, err := client.Repositories.GetBranch(ctx, cfg.GithubRepoOwner, cfg.GithubRepo, "main")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rate: %#v\n", resp.Rate)

	if !resp.TokenExpiration.IsZero() {
		log.Printf("Token Expiration: %v\n", resp.TokenExpiration)
	}

	fmt.Printf("\n%v\n", github.Stringify(user))
}
