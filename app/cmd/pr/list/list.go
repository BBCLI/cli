package list

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "list pull requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		state := "OPEN"
		params := bbclient.GetPullrequestsSelectedUserParams{
			State: (*bbclient.GetPullrequestsSelectedUserParamsState)(&state),
		}
		res, err := bbclient.BbClient.GetPullrequestsSelectedUserWithResponse(context.TODO(), "user-H-u20S8xQNiGvn-hnSLGqQ", &params)
		if err != nil {
			return err
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
