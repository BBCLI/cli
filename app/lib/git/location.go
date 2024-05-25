package git

import (
	"os/exec"
	"strings"
)

func GetGitRemoteDetails() (workspace string, repo string, err error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	output, err := cmd.Output()
	if err != nil {
		return "", "", err
	}

	// The remote URL could be in either of these formats:
	// SSH: git@github.com:<workspace>/<repo>.git
	// HTTPS: https://github.com/<workspace>/<repo>.git
	// Split the URL by either ':' (for SSH) or '/' (for HTTPS)
	var parts []string
	if strings.Contains(string(output), "https://") {
		parts = strings.Split(string(output), "/")
		if len(parts) < 5 {
			return "", "", nil
		}
		workspace = parts[3]
		repo = strings.TrimSuffix(strings.TrimSpace(parts[4]), ".git")
	} else {
		parts = strings.Split(string(output), ":")
		if len(parts) < 2 {
			return "", "", nil
		}
		workspaceAndRepo := strings.Split(parts[1], "/")
		if len(workspaceAndRepo) < 2 {
			return "", "", nil
		}
		workspace = workspaceAndRepo[0]
		repo = strings.TrimSuffix(strings.TrimSpace(workspaceAndRepo[1]), ".git")
	}

	return workspace, repo, nil
}
