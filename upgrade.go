package main

import (
	"github.com/tcnksm/go-latest"
)

const version = "0.1.0"

// NeedsUpgrade check if verto desktop is updated or not
func NeedsUpgrade() bool {
	githubTag := &latest.GithubTag{
		Owner:      "useverto",
		Repository: "desktop",
	}

	res, _ := latest.Check(githubTag, version)
	return res.Outdated
}
