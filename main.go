package main

import (
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/http"
	"context"
	"fmt"
	"os"
)

// Model
type RepoInfo struct {
	FullName      string
	StarsCount    int
	ForksCount    int
	LastUpdatedBy string
}


func main() {
	// Github Auth
	if os.Getenv("GITHUB_API_KEY") == "" {
		fmt.Printf("Need to set GITHUB_API_KEY")
		os.Exit(1)
	}
	context := context.Background()
	tokenService := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_API_KEY")},
	)
	tokenClient := oauth2.NewClient(context, tokenService)
	client := github.NewClient(tokenClient)

	// Take argument
	cmd := os.Args[1]
	if cmd == "info" {
		fmt.Printf("Here is info\n")
		org := os.Args[2]
		repoName := os.Args[3]

		repo, _, err := client.Repositories.Get(context, org, repoName)

		if err != nil {
			fmt.Printf("Problem in getting repository information %v\n", err)
			os.Exit(1)
		}

		pack := &RepoInfo{
			FullName: *repo.FullName,
			ForksCount: *repo.ForksCount,
			StarsCount: *repo.StargazersCount,
		}

		fmt.Printf("%+v\n", pack)

	} else if cmd == "delete" {
		fmt.Printf("Here is delete repo\n")
		org := os.Args[2]
		repoName := os.Args[3]

		_, err := client.Repositories.Delete(context, org, repoName)
		if err != nil {
			fmt.Printf("Repositories.Delete() returned error: %v", err)
			os.Exit(1)
		}

	} else if cmd == "create" {
		fmt.Printf("Here is create repo\n")
		org := os.Args[2]
		repoName := os.Args[3]

		_, resp, err := client.Repositories.Get(context, org, repoName)
		if err != nil {
			if resp.StatusCode != http.StatusNotFound {
				fmt.Printf("Repo already existed, exit!")
				os.Exit(1)
			}
		}

		_, _, err = client.Repositories.Create(context, "", &github.Repository{
			Name:     github.String(repoName),
			AutoInit: github.Bool(false),
			Private:  github.Bool(false),
		})
		if err != nil {
			fmt.Printf("Repositories.Create() returned error: %v", err)
			os.Exit(1)
		}
	} else if cmd == "user-list-repo" {
		fmt.Printf("Here is user-list-repo\n")
		user := os.Args[2]

		opt := &github.RepositoryListOptions{
			ListOptions: github.ListOptions{PerPage: 10},
		}
		// get all pages of results
		var allRepos []*github.Repository
		for {
			repos, resp, err := client.Repositories.List(context, user, opt)
			if err != nil {
				fmt.Printf("Repositories.List() return error: %v", err)
				os.Exit(1)
			}
			allRepos = append(allRepos, repos...)
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}

		for _, repo := range allRepos {
			pack := &RepoInfo{
				FullName: *repo.FullName,
				ForksCount: *repo.ForksCount,
				StarsCount: *repo.StargazersCount,
			}

			fmt.Printf("%+v\n", pack)
		}

	} else if cmd == "org-list-repo" {
		fmt.Printf("Here is org-list-repo\n")
		user := os.Args[2]

		opt := &github.RepositoryListByOrgOptions{
			ListOptions: github.ListOptions{PerPage: 10},
		}
		// get all pages of results
		var allRepos []*github.Repository
		for {
			repos, resp, err := client.Repositories.ListByOrg(context, user, opt)
			if err != nil {
				fmt.Printf("Repositories.ListByOrg() return error: %v", err)
				os.Exit(1)
			}
			allRepos = append(allRepos, repos...)
			if resp.NextPage == 0 {
				break
			}
			opt.Page = resp.NextPage
		}

		for _, repo := range allRepos {
			pack := &RepoInfo{
				FullName: *repo.FullName,
				ForksCount: *repo.ForksCount,
				StarsCount: *repo.StargazersCount,
			}

			fmt.Printf("%+v\n", pack)
		}
	} else {
		fmt.Printf("No such command: %s\n", cmd)
		os.Exit(1)
	}
}
