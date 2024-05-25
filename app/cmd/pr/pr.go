package pr

import (
	"cli/app/cmd/pr/approve"
	"cli/app/cmd/pr/comment"
	"cli/app/cmd/pr/create"
	"cli/app/cmd/pr/diff"
	"cli/app/cmd/pr/list"
	"cli/app/cmd/pr/merge"
	"cli/app/cmd/pr/show"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "pr",
	Short: "PR related stuff",
}

func init() {
	Cmd.AddCommand(list.Cmd)
	Cmd.AddCommand(create.Cmd)
	Cmd.AddCommand(approve.Cmd)
	Cmd.AddCommand(merge.Cmd)
	Cmd.AddCommand(diff.Cmd)
	Cmd.AddCommand(show.Cmd)
	Cmd.AddCommand(comment.Cmd)
}
