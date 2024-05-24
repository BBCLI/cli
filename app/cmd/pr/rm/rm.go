package rm

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "rm",
	Short: "delete pull requests",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Prs index")

		return nil
	},
}
