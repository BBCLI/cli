package me

import (
	"fmt"

	"github.com/spf13/cobra"

	"cli/app/cmd/me/pr"
)

var Cmd = &cobra.Command{
	Use:   "me",
	Short: "Me related commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Prs index")
	},
}

func init() {
	Cmd.AddCommand(pr.Cmd)
}
