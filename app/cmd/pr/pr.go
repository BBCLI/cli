package pr

import (
	"cli/app/cmd/pr/list"
	"fmt"
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
}
