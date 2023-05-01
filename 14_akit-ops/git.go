// It has functions to clone a repo and detect any changes in the repo
package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
)

// clone the provided repo
func cloneRepo(url, repoPath string) *git.Repository {
	r, err := git.PlainClone(repoPath, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	checkIfError(err)
	return r
}

// Detect any changes in the repo
func detectChanges(r *git.Repository) bool {

	isChanged := false

	ref, err := r.Head()

	checkIfError(err)

	currentHashId := ref.Hash()

	// Get the working directory for the repository
	w, err := r.Worktree()

	checkIfError(err)

	w.Pull(&git.PullOptions{RemoteName: "origin"})

	ref, err = r.Head()

	checkIfError(err)
	newHashId := ref.Hash()

	if currentHashId != newHashId {
		fmt.Println(string("\033[32m"), "Change detected...Time to deploy", string("\033[0m"))
		isChanged = true
	} else {
		fmt.Println(string("\033[31m"), "No change detected...", string("\033[0m"))
	}

	return isChanged

}
