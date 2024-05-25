package resolve

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	"cli/app/lib/git"
)

var Cmd = &cobra.Command{
	Use:   "resolve",
	Short: "resolve comment on PR",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		prId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		commentId, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		workspace, repo, err := git.GetGitRemoteDetails()
		if err != nil {
			return err
		}

		res, err := bbclient.BbClient.PostRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsCommentIdResolveWithResponse(context.TODO(), workspace, repo, prId, commentId)
		if err != nil || res.StatusCode() != http.StatusOK {
			return err
		}

		fmt.Println("comment resolved")
		return nil
	},
}
