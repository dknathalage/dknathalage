package gitutil

import (
	"context"
	"log"

	"github.com/google/go-github/v66/github"
)

func GetDefaultBranch(client *github.Client, ctx context.Context, owner, repo string) string {
	branch, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	return branch.GetDefaultBranch()
}

func IsDefaultBranch(client *github.Client, ctx context.Context, owner, repo, ref string) bool {
	defaultBranch := GetDefaultBranch(client, ctx, owner, repo)
	return ref == defaultBranch
}
