package pr

import (
	"fmt"

	"cli/app/cmd/pr/create"
	"cli/app/cmd/pr/list"
	"cli/app/cmd/pr/rm"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "pr",
	Short: "PR related stuff",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Prs index")

		return nil
	},
}

func init() {
	Cmd.AddCommand(list.Cmd)
	Cmd.AddCommand(rm.Cmd)
	Cmd.AddCommand(create.Cmd)
}
