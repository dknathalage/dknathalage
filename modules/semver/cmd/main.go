package main

import (
	"context"
	"errors"

	"github.com/dknathalage/dknathalage/modules/semver/gitutil"
	"github.com/google/go-github/v66/github"
)

func validateCfg(cfg *gitutil.Config, client *github.Client, ctx *context.Context) error {
	if cfg.GithubRefType != "branch" {
		return errors.New("GITHUB_REF_TYPE must be branch")
	}

	if !gitutil.IsDefaultBranch(client, ctx, cfg.GithubRepoOwner, cfg.GithubRepository, cfg.GithubRef) {
		return errors.New("GITHUB_REF must be the default branch")
	}

	return nil
}

func main() {
	cfg := gitutil.LoadConfig()
	ctx := context.Background()
	client := github.NewClient(nil).WithAuthToken(string(cfg.GithubToken))

	err := validateCfg(&cfg, client, &ctx)
	if err != nil {
		panic(err)
	}
}
