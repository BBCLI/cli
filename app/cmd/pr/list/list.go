package list

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	"cli/app/lib/fetch"
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
			State: (*bbclient.GetRepositoriesWorkspaceRepoSlugPullrequestsParamsState)(&state),
		}
		workspace, repo, err := git.GetGitRemoteDetails()
		if err != nil {
			log.Fatal("No Workspace Or Repo Found! Please Check If are in a Repo that has a remote origin in bbc")
		}

		res, err := bbclient.BbClient.GetRepositoriesWorkspaceRepoSlugPullrequestsWithResponse(context.TODO(), workspace, repo, &params)
		if err != nil || res.StatusCode() != http.StatusOK {
			log.Fatal(fmt.Sprintf("Error From The Bitbucket Cloud Api! Status: %v", res.StatusCode()))
		}
		prs := *res.JSON200.Values
		nextUrl := res.JSON200.Next

		for nextUrl != nil {
			pageRes, err := fetch.FetchWithAuth(*nextUrl)
			if pageRes.StatusCode != 200 {
				log.Fatal(fmt.Sprintf("Error Fetching Paginated PRs %v", pageRes.StatusCode))
				return nil
			}
			if err != nil {
				return err
			}
			pageData, err := bbclient.ParseGetRepositoriesWorkspaceRepoSlugPullrequestsResponse(pageRes)
			if err != nil {
				return err
			}
			nextUrl = pageData.JSON200.Next
			prs = append(prs, *pageData.JSON200.Values...)
		}
		err = format.Prs(&prs)
		if err != nil {
			return err
		}
		return nil
	},
}
