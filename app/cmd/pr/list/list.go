package list

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
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
		if err != nil {
			return err
		}
		prs := *res.JSON200.Values

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)
		fmt.Fprintln(w, fmt.Sprintf("%v\t  %s\t %s\t %s", "id", "title", "commentCount", "repository"))
		for i := 0; i < len(prs); i++ {
			pr := (prs)[i]
			_, err := fmt.Fprintln(w, fmt.Sprintf("%v\t  %s\t %v\t %s ", *pr.Id, *pr.Title, *pr.CommentCount, *pr.Source.Repository.Name))
			if err != nil {
				return err
			}
		}
		err = w.Flush()
		if err != nil {
			return err
		}

		return nil
	},
}
