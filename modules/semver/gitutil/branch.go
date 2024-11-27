package gitutil

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/go-github/v66/github"
)

func GetDefaultBranch(client *github.Client, ctx context.Context, repo string) string {
	index := strings.Index(repo, "/")
	if index == -1 {
		log.Fatal(fmt.Errorf("invalid repo format: %s", repo))
	}

	owner := repo[:index]
	repo = repo[index+1:]
	branch, _, err := client.Repositories.Get(ctx, owner, repo)
	if err != nil {
		log.Fatal(err)
	}

	return branch.GetDefaultBranch()
}

func IsDefaultBranch(ctx *context.Context, client *github.Client, repo, ref string) bool {
	defaultBranch := GetDefaultBranch(client, *ctx, repo)
	return ref == "refs/heads/"+defaultBranch
}
