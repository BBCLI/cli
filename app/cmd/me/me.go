package me

import (
	"github.com/spf13/cobra"

	"cli/app/cmd/me/pr"
)

var Cmd = &cobra.Command{
	Use:   "me",
	Short: "Me related commands",
}

func init() {
	Cmd.AddCommand(pr.Cmd)
}
