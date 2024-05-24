package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Create a pull request",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Prs index")
	},
}
