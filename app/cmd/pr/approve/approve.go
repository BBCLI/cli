package approve

import (
	"context"
	"errors"
	"strconv"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	"cli/app/lib/git"
)

var Cmd = &cobra.Command{
	Use:   "approve",
	Short: "List pull requests",
	Long:  "usage: bbcli pr approve <pullrequest-id>",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		prId, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		workspace, repo, err := git.GetGitRemoteDetails()
		if err != nil {
			return err
		}

		res, err := bbclient.BbClient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveWithResponse(context.TODO(), workspace, repo, prId)
		if err != nil || res.StatusCode() != 200 {
			return errors.New("an error occurred")
		}
		return nil
	},
}
