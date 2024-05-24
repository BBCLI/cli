package list

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "list pull requests",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("List of Prs")
	},
}
