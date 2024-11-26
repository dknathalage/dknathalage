package main

import (
	"context"
	"log"

	"github.com/google/go-github/v66/github"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
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
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repo, _, err := client.Repositories.Get(ctx, "dknathalage", "dknathalage")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Repository: %s", *repo.FullName)
}
