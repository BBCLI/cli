package create

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	"cli/app/lib/git"
	"cli/app/lib/user"
)

var Cmd = &cobra.Command{
	Use:     "create",
	Short:   "create a pull request",
	Long:    "usage: pr create <origin-branch> <destination-branch> | pr create <destination-branch>",
	Args:    cobra.RangeArgs(1, 2),
	Aliases: []string{"new"},
	RunE: func(cmd *cobra.Command, args []string) error {
		var origin, dest string
		if len(args) == 1 {
			dest = args[0]
			currentBranch, err := git.GetCurrentBranch()
			if err != nil {
				fmt.Println("Error Reading current Branch")
				return nil
			}
			origin = currentBranch
		} else {
			origin = args[0]
			dest = args[1]
		}
		currentUser, err := user.CurrentUser()
		if err != nil {
			log.Fatal("Error while fetching user")
			return nil
		}

		workspace, repo, err := git.GetGitRemoteDetails()
		if err != nil {
			return err
		}

		defaultTitle := origin + " to " + dest
		fmt.Printf("Please enter PR title ('%s'): ", defaultTitle)
		in := bufio.NewReader(os.Stdin)
		title, err := in.ReadString('\n')
		if err != nil {
			return errors.New("error Reading your title")
		}
		title = strings.TrimSpace(title)
		fmt.Println(title)
		if title == "" {
			title = defaultTitle
		}
		//default reviewers(repo)
		fmt.Printf("Do you want to add the default reviewers from the repo(y/n): ")
		in = bufio.NewReader(os.Stdin)
		desiscion, err := in.ReadString('\n')
		desiscion = strings.TrimSpace(desiscion)
		var reviewers []bbclient.Account
		if desiscion == "y" || desiscion == "yes" {
			res, err := bbclient.BbClient.GetRepositoriesWorkspaceRepoSlugDefaultReviewersWithResponse(context.TODO(),
				workspace, repo)
			if err != nil || res.StatusCode() != 200 {
				return errors.New("Error getting Reviewers from bitbucket url")
			}
			reviewers = *res.JSON200.Values
		}
		var reviewersWithoutCurrentUser []bbclient.Account
		reviewersWithoutCurrentUser = make([]bbclient.Account, 0)
		//remove user
		for i := range reviewers {
			if *reviewers[i].Uuid != *currentUser.Uuid {
				reviewersWithoutCurrentUser = append(reviewersWithoutCurrentUser, reviewers[i])
			}
		}

		req := bbclient.PostRepositoriesWorkspaceRepoSlugPullrequestsJSONRequestBody{
			Title:     &title,
			Reviewers: &reviewersWithoutCurrentUser,
			Destination: &bbclient.PullrequestEndpoint{
				Branch: &struct {
					DefaultMergeStrategy *string                                              `json:"default_merge_strategy,omitempty"`
					MergeStrategies      *[]bbclient.PullrequestEndpointBranchMergeStrategies `json:"merge_strategies,omitempty"`
					Name                 *string                                              `json:"name,omitempty"`
				}{
					Name: &dest,
				},
			},
			Source: &bbclient.PullrequestEndpoint{
				Branch: &struct {
					DefaultMergeStrategy *string                                              `json:"default_merge_strategy,omitempty"`
					MergeStrategies      *[]bbclient.PullrequestEndpointBranchMergeStrategies `json:"merge_strategies,omitempty"`
					Name                 *string                                              `json:"name,omitempty"`
				}{
					Name: &origin,
				},
			},
		}

		res, err := bbclient.BbClient.PostRepositoriesWorkspaceRepoSlugPullrequestsWithResponse(context.TODO(), workspace, repo, req)

		if err != nil || res.StatusCode() != 201 {
			jsonBytes, _ := json.Marshal(res.JSON400)
			fmt.Println(string(jsonBytes))
			return errors.New(fmt.Sprintf("Error while creating requesting the bbc api! status: %v", res.StatusCode()))
		}
		fmt.Printf("PR Created! Link: %s\n", *res.JSON201.Links.Html.Href)
		return nil
	},
}
