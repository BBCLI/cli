package rm

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "rm",
	Short:   "Delete pull requests",
	Aliases: []string{"delete"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Prs index")
	},
}
