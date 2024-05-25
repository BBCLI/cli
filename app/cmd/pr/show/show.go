package show

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path"
	"strconv"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	"cli/app/lib/git"
)

var Cmd = &cobra.Command{
	Use:   "show",
	Short: "Show pull request details",
	Long:  "usage: bbc pr show <pullrequest-id>",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		prId, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		prId = 1197
		workspace, repo, err := git.GetGitRemoteDetails()
		workspace = "check24"
		repo = "tippspiel"
		if err != nil {
			return err
		}

		sort := "-created_on"
		response, err := bbclient.BbClient.GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesWithResponse(context.TODO(), workspace, repo, prId, &bbclient.GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdStatusesParams{
			Sort: &sort,
		})
		if err != nil || response.StatusCode() != http.StatusOK {
			return errors.New("error on bbc api")
		}

		latestCommits := *response.JSON200.Values
		if len(latestCommits) == 0 {
			return errors.New("no commit statuses found")
		}
		symbol := ""
		commit := latestCommits[0]
		commitState := *commit.State
		if commitState == bbclient.CommitstatusStateSUCCESSFUL {
			symbol = "✅ "
		}
		if commitState == bbclient.CommitstatusStateFAILED {
			symbol = "❌ "
		}
		if commitState == bbclient.CommitstatusStateINPROGRESS {
			symbol = "🚧 "
		}

		prResponse, err := bbclient.BbClient.GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdWithResponse(context.TODO(), workspace, repo, prId)
		if err != nil || prResponse.StatusCode() != http.StatusOK {
			return errors.New("error on bbc api")
		}

		commentResponse, err := bbclient.BbClient.GetRepositoriesWorkspaceRepoSlugPullrequestsPullRequestIdCommentsWithResponse(context.TODO(), workspace, repo, prId)
		if err != nil || commentResponse.StatusCode() != http.StatusOK {
			return errors.New("error on bbc api")
		}

		fmt.Printf("from %s to %s\n", *prResponse.JSON200.Source.Branch.Name, *prResponse.JSON200.Destination.Branch.Name)
		fmt.Printf("Title: %s\n", *prResponse.JSON200.Title)
		fmt.Printf("Created By: %s\n", *prResponse.JSON200.Author.DisplayName)
		if *prResponse.JSON200.Summary.Raw != "" {
			fmt.Printf("\n%s\n\n", *prResponse.JSON200.Summary.Raw)
		}

		fmt.Printf("%s: %s %s \n", *commit.Name, commitState, symbol)

		fmt.Printf("latest commit created at: %s \n\n", *prResponse.JSON200.UpdatedOn)

		fmt.Printf("Comments: \n")

		for _, comment := range *commentResponse.JSON200.Values {
			isCommentOnFile := false
			inline := comment.Inline
			line := 0
			if inline != nil {
				isCommentOnFile = true
				line = *inline.To
			}
			resolvedText := "resolved"
			if comment.Resolution == nil {
				resolvedText = "pending"
			}

			if isCommentOnFile {
				repoPath, err := git.GetGitAbsolutePath()
				if err != nil {
					return err
				}
				path := path.Join(repoPath, inline.Path)
				fmt.Printf("%s commented on %s:%v (%s) :\n", *comment.User.DisplayName, path, line, resolvedText)
				fmt.Printf("%s\n", *comment.Content.Raw)
				continue
			}
			fmt.Printf("%s commented (%s) :\n", *comment.User.DisplayName, resolvedText)
			fmt.Printf("%s\n", *comment.Content.Raw)
			fmt.Printf("-------------------------------------------\n")
		}
		return nil
	},
}
