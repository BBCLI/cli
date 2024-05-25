package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetGitRemoteDetails() (workspace string, repo string, err error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	output, err := cmd.Output()
	if err != nil || !strings.Contains(string(output), "bitbucket.org") {
		fmt.Println("Looks like you are not in a git repo, or it doesn't belong to bitbucket cloud(not bitbucket.org in url)")
		workspace, repo = promptForWorkspaceAndRepo()
		return workspace, repo, nil
	}

	// The remote URL could be in either of these formats:
	// SSH: git@bitbucket.org:<workspace>/<repo>.git
	// HTTPS: https://bitbucket.org/<workspace>/<repo>.git
	// Split the URL by either ':' (for SSH) or '/' (for HTTPS)
	var parts []string
	if strings.Contains(string(output), "https://") {
		parts = strings.Split(string(output), "/")
		if len(parts) < 5 {
			workspace, repo = promptForWorkspaceAndRepo()
			return workspace, repo, nil
		}
		workspace = parts[3]
		repo = strings.TrimSuffix(strings.TrimSpace(parts[4]), ".git")
	} else {
		parts = strings.Split(string(output), ":")
		if len(parts) < 2 {
			workspace, repo = promptForWorkspaceAndRepo()
			return workspace, repo, nil
		}
		workspaceAndRepo := strings.Split(parts[1], "/")
		if len(workspaceAndRepo) < 2 {
			workspace, repo = promptForWorkspaceAndRepo()
			return workspace, repo, nil
		}
		workspace = workspaceAndRepo[0]
		repo = strings.TrimSuffix(strings.TrimSpace(workspaceAndRepo[1]), ".git")
	}

	return workspace, repo, nil
}

func GetGitAbsolutePath() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	ret, _ := strings.CutSuffix(string(output), "\n")
	return ret, nil
}

func promptForWorkspaceAndRepo() (string, string) {
	var workspace, repo string
	fmt.Print("Enter workspace: ")
	fmt.Scanln(&workspace)
	fmt.Print("Enter repo: ")
	fmt.Scanln(&repo)
	return workspace, repo
}
