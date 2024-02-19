package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Source GitLab repository URL
	// can do beeter like loop from config file
	gitlabRepo := "https://gitlab.com/yourusername/yourrepository.git"

	// Destination GitHub repository URL
	githubRepo := "https://github.com/yourusername/yourrepository.git"

	// Clone the GitLab repository
	cmd := exec.Command("git", "clone", "--bare", gitlabRepo)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error cloning GitLab repository:", err)
		os.Exit(1)
	}

	// Change directory to the cloned repository
	if err := os.Chdir("yourrepository.git"); err != nil {
		fmt.Println("Error changing directory:", err)
		os.Exit(1)
	}

	// Push all branches to GitHub
	pushBranchesCmd := exec.Command("git", "push", "--mirror", githubRepo)
	if err := pushBranchesCmd.Run(); err != nil {
		fmt.Println("Error pushing branches to GitHub:", err)
		os.Exit(1)
	}

	// Push all tags to GitHub
	pushTagsCmd := exec.Command("git", "push", "--tags", githubRepo)
	if err := pushTagsCmd.Run(); err != nil {
		fmt.Println("Error pushing tags to GitHub:", err)
		os.Exit(1)
	}

	// Clean up
	if err := os.Chdir(".."); err != nil {
		fmt.Println("Error changing directory:", err)
		os.Exit(1)
	}
	if err := os.RemoveAll("yourrepository.git"); err != nil {
		fmt.Println("Error removing directory:", err)
		os.Exit(1)
	}

	fmt.Println("Migration completed successfully!")
}

func clone(string remoteRepo = nil, string destRepo = nil ) {
	// only clone repo
}
