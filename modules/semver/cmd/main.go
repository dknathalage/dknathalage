package main

import (
	"context"
	"fmt"
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

	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	for {
		repos, resp, err := client.Repositories.List(ctx, "", opt)
		if err != nil {
			log.Fatalf("Error fetching repositories: %v", err)
		}

		for _, repo := range repos {
			fmt.Printf("Name: %s, URL: %s\n", *repo.Name, *repo.HTMLURL)
		}

		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}
}
