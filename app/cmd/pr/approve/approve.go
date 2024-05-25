package approve

import (
	"context"
	"errors"
	"strconv"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
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
		//TODO dynamic workspace and repo
		res, err := bbclient.BbClient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdApproveWithResponse(context.TODO(), "check24", "tippspiel", prId)
		if err != nil || res.StatusCode() != 200 {
			return errors.New("an error occurred")
		}
		return nil
	},
}
