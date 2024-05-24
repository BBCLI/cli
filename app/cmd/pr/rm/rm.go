package rm

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "rm",
	Short: "delete pull requests",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Prs index")
	},
}
