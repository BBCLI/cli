package comment

import (
	"github.com/spf13/cobra"

	"cli/app/cmd/pr/comment/create"
	"cli/app/cmd/pr/comment/reply"
	"cli/app/cmd/pr/comment/resolve"
)

var Cmd = &cobra.Command{
	Use:   "comment",
	Short: "Pr comment related commands",
}

func init() {
	Cmd.AddCommand(reply.Cmd)
	Cmd.AddCommand(create.Cmd)
	Cmd.AddCommand(resolve.Cmd)
}
