package list

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List PRs",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("List of Prs")

		return nil
	},
}
