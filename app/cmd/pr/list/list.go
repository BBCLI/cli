package list

import (
	"context"
	"fmt"
	"log"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"

	"cli/app/bbclient"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "list pull requests",
	Run: func(cmd *cobra.Command, args []string) {
		state := "OPEN"
		params := bbclient.GetPullrequestsSelectedUserParams{
			State: (*bbclient.GetPullrequestsSelectedUserParamsState)(&state),
		}
		res, err := bbclient.BbClient.GetPullrequestsSelectedUserWithResponse(context.TODO(), "user-H-u20S8xQNiGvn-hnSLGqQ", &params)
		if err != nil {
			log.Fatal(err)
		}
		prs := *res.JSON200.Values

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)
		for i := 0; i < len(prs); i++ {
			pr := (prs)[i]
			fmt.Fprintln(w, fmt.Sprintf("%v\t  %s", *pr.Id, *pr.Title))
		}
		w.Flush()
		//marshal, err := json.Marshal(res.JSON200)
		//log.Printf("here: %s\n", marshal)
	},
}
