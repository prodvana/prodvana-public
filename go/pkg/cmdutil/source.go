package cmdutil

import (
	"fmt"
	"os"
)

type CIEnvironment string

const (
	CircleciEnvironment      CIEnvironment = "circleci"
	GitHubActionsEnvironment CIEnvironment = "gha"
)

func detectCI() CIEnvironment {
	if os.Getenv("CIRCLECI") == "true" {
		return CircleciEnvironment
	}
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		return GitHubActionsEnvironment
	}
	return ""
}

type CIDefaults struct {
	Environment   CIEnvironment
	Repository    string
	Commit        string
	BaseDirectory string
}

func DetectCIDefaults() *CIDefaults {
	switch detectCI() {
	case CircleciEnvironment:
		// https://circleci.com/docs/variables/#built-in-environment-variables
		return &CIDefaults{
			Environment:   CircleciEnvironment,
			Commit:        os.Getenv("CIRCLE_SHA1"),
			Repository:    os.Getenv("CIRCLE_REPOSITORY_URL"),
			BaseDirectory: os.Getenv("CIRCLE_WORKING_DIRECTORY"),
		}
	case GitHubActionsEnvironment:
		// https://docs.github.com/en/actions/learn-github-actions/variables#default-environment-variables
		return &CIDefaults{
			Environment:   GitHubActionsEnvironment,
			Commit:        os.Getenv("GITHUB_SHA"),
			Repository:    fmt.Sprintf("%s/%s", os.Getenv("GITHUB_SERVER_URL"), os.Getenv("GITHUB_REPOSITORY")),
			BaseDirectory: os.Getenv("GITHUB_WORKSPACE"),
		}
	default:
		return nil
	}
}
