package pr

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
	config2 "cli/app/lib/config"
	"cli/app/lib/format"
)

var Cmd = &cobra.Command{
	Use:   "pr",
	Short: "List your pull requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		state := bbclient.OPEN
		params := bbclient.GetPullrequestsSelectedUserParams{
			State: (*bbclient.GetPullrequestsSelectedUserParamsState)(&state),
		}
		config, err := config2.GetConfig()
		if err != nil {
			return err
		}
		if config.Authorization.Username == "" || config.Authorization.Password == "" {
			return fmt.Errorf("please run 'bbcli init' to initialize your Bitbucket Cloud CLI")
		}
		res, err := bbclient.BbClient.GetPullrequestsSelectedUserWithResponse(context.TODO(), config.Authorization.Username, &params)
		if err != nil {
			return err
		}
		if res.JSON200 == nil {
			return fmt.Errorf("couldn't fetch PRs")
		}

		err = format.Prs(res.JSON200.Values)
		if err != nil {
			return err
		}

		return nil
	},
}
