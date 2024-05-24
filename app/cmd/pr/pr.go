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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Prs index")
	},
}

func init() {
	Cmd.AddCommand(list.Cmd)
	Cmd.AddCommand(rm.Cmd)
	Cmd.AddCommand(create.Cmd)
}
