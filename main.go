package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/golang/glog"
	"github.com/google/go-github/github"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		glog.Fatalf("Usage: go-detect <github_repo>")
	}
	parts := strings.Split(flag.Arg(0), "/")
	if len(parts) < 2 {
		glog.Fatalf("Invalid URL %s", flag.Arg(0))
	}
	client := github.NewClient(nil)
	repo, _, err := client.Repositories.Get(parts[len(parts)-2], parts[len(parts)-1])
	if err != nil {
		glog.Fatalf("Github error: %v", err)
	}
	fmt.Printf("Language: %s\n", *repo.Language)
}
