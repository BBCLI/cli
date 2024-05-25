package diff

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	"cli/app/lib/git"
)

var Cmd = &cobra.Command{
	Use:     "diff",
	Short:   "show diff for pull request",
	Aliases: []string{"ls"},
	Args:    cobra.ExactArgs(1),
	Long:    "usage: bbc pr diff <pullrequest-id>",
	RunE: func(cmd *cobra.Command, args []string) error {
		workspace, repo, err := git.GetGitRemoteDetails()
		if err != nil {
			log.Fatal("No Workspace Or Repo Found! Please Check If are in a Repo that has a remote origin in bbc")
		}
		prId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		diff, err := bbclient.BbClient.GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdDiffWithResponse(context.TODO(), workspace, repo, prId)
		if err != nil {
			return err
		}
		if diff.StatusCode() != 200 {
			return errors.New("couldn't fetch PR")
		}
		log.Println(string(diff.Body))
		return nil
	},
}
