package list

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	"cli/app/lib/format"
	"cli/app/lib/git"
)

var Cmd = &cobra.Command{
	Use:     "list",
	Short:   "list pull requests",
	Aliases: []string{"ls"},
	RunE: func(cmd *cobra.Command, args []string) error {
		state := bbclient.OPEN
		params := bbclient.GetRepositoriesWorkspaceRepoSlugPullrequestsParams{
			State: &state,
		}
		workspace, repo, err := git.GetGitRemoteDetails()
		if err != nil {
			log.Fatal("No Workspace Or Repo Found! Please Check If are in a Repo that has a remote origin in bbc")
		}
		res, err := bbclient.BbClient.GetRepositoriesWorkspaceRepoSlugPullrequestsWithResponse(context.TODO(), workspace, repo, &params)
		if err != nil || res.StatusCode() != http.StatusOK {
			log.Fatal(fmt.Sprintf("Error From The Bitbucket Cloud Api! Status: %v", res.StatusCode()))
			return nil
		}

		err = format.Prs(res.JSON200.Values)
		if err != nil {
			return err
		}

		return nil
	},
}
