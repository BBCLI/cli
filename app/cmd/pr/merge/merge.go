package merge

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	"cli/app/lib/git"
)

var Cmd = &cobra.Command{
	Use:   "merge",
	Short: "merge pull request",
	Long:  "usage: bbcli pr merge <pullrequest-id>",
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

		res, err := bbclient.BbClient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeWithResponse(context.TODO(), workspace, repo, prId, &bbclient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeParams{}, bbclient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdMergeJSONRequestBody{})
		if err != nil || res.StatusCode() != 200 {
			return errors.New("an error occurred")
		}

		fmt.Println("pr merged successfully")

		return nil
	},
}
