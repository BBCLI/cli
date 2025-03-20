package open

import (
	"context"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	"cli/app/lib/git"
)

var Cmd = &cobra.Command{
	Use:   "open",
	Short: "open a link to pr",
	Long:  "usage: bbc pr open <pullrequest-id>",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		prId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		workspace, repo, err := git.GetGitRemoteDetails()
		if err != nil {
			return err
		}

		res, err := bbclient.BbClient.GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdWithResponse(context.TODO(), workspace, repo, prId)
		if err != nil {
			return err
		}
		err = openLink(res.JSON200.Links.Html.Href)
		if err != nil {
			return err
		}
		return nil
	},
}

func openLink(url *string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", *url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", *url)
	default: // assume Linux or other Unix-like OS
		cmd = exec.Command("xdg-open", *url)
	}

	return cmd.Start()
}
