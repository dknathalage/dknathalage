package main

import (
	"context"
	"log"

	"github.com/dknathalage/modules/semver/gitutil"
	"github.com/google/go-github/v66/github"
)

func main() {
	cfg := gitutil.LoadConfig()
	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(string(cfg.GithubToken))

	if cfg.GithubRefType != "branch" {
		log.Fatal("GITHUB_REF_TYPE must be branch")
	}

	if !gitutil.IsDefaultBranch(client, ctx, cfg.GithubRepoOwner, cfg.GithubRepository, cfg.GithubRef) {
		log.Fatal("GITHUB_REF must be the default branch")
	}
}
