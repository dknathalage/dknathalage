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
		&oauth2.Token{AccessToken: cfg.GithubToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	releases, _, err := client.Repositories.ListReleases(ctx, "dknathalage", "dknagalage", nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, release := range releases {
		log.Printf("Name: %s, Tag: %s", release.GetName(), release.GetTagName())
	}
}
