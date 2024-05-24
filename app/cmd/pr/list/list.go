package list

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	config2 "cli/app/lib/config"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "list pull requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		state := "OPEN"
		params := bbclient.GetPullrequestsSelectedUserParams{
			State: (*bbclient.GetPullrequestsSelectedUserParamsState)(&state),
		}
		config, err := config2.GetConfig()
		if err != nil {
			return err
		}
		if config.Authorization.Username == "" || config.Authorization.Password == "" {
			return fmt.Errorf("please run 'bbc init' to initialize your Bitbucket Cloud CLI")
		}
		res, err := bbclient.BbClient.GetPullrequestsSelectedUserWithResponse(context.TODO(), config.Authorization.Username, &params)
		if err != nil {
			return err
		}
		if res.JSON200 == nil {
			return fmt.Errorf("couldn't fetch PRs")
		}
		prs := *res.JSON200.Values

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)
		for i := 0; i < len(prs); i++ {
			pr := (prs)[i]
			_, err := fmt.Fprintln(w, fmt.Sprintf("%v\t  %s", *pr.Id, *pr.Title))
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
