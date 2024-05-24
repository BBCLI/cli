package create

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
)

var Cmd = &cobra.Command{
	Use:     "create",
	Short:   "create a pull request",
	Long:    "usage: pr create origin-branch destination-branch",
	Args:    cobra.ExactArgs(2),
	Aliases: []string{"new"},
	RunE: func(cmd *cobra.Command, args []string) error {
		origin := args[0]
		dest := args[1]

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

		req := bbclient.PostRepositoriesWorkspaceRepoSlugPullrequestsJSONRequestBody{
			Title: &title,
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
		res, err := bbclient.BbClient.PostRepositoriesWorkspaceRepoSlugPullrequestsWithResponse(context.TODO(), "check24", "tippspiel", req)
		if err != nil || res.StatusCode() != 201 {
			return errors.New("an error occurred")
		}

		return nil
	},
}
