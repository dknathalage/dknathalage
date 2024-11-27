package gitutil

import (
	"context"
	"log"
	"strings"

	"github.com/google/go-github/v66/github"
)

func GetDefaultBranch(client *github.Client, ctx context.Context, repo string) string {
	owner := repo[:strings.Index(repo, "/")]
	repo = repo[strings.Index(repo, "/")+1:]
	branch, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	return branch.GetDefaultBranch()
}

func IsDefaultBranch(client *github.Client, ctx *context.Context, owner, repo, ref string) bool {
	defaultBranch := GetDefaultBranch(client, *ctx, repo)
	return ref == "refs/heads/"+defaultBranch
}
