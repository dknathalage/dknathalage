package gitutil

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	GithubToken      string `envconfig:"GITHUB_TOKEN" required:"true"`
	GithubRepoOwner  string `envconfig:"GITHUB_REPO_OWNER"`
	GithubRepository string `envconfig:"GITHUB_REPOSITORY"`
	GithubRefType    string `envconfig:"GITHUB_REF_TYPE"`
	GithubRef        string `envconfig:"GITHUB_REF"`
}

func LoadConfig() Config {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
