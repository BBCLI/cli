package create

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "create",
	Short: "Create a pull request",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("Prs index")

		return nil
	},
}
