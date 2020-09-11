package main

import (
	"fmt"

	"github.com/tcnksm/go-latest"
)

const version = "0.1.0"

// NeedsUpgrade check if verto desktop is updated or not
func NeedsUpgrade() bool {
	githubTag := &latest.GithubTag{
		Owner:      "useverto",
		Repository: "desktop",
	}

	res, err := latest.Check(githubTag, version)
	if err != nil {
		fmt.Println("Unable to compare github version")
		return false
	}
	return res.Outdated
}
