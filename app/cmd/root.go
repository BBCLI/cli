package cmd

import (
	initialize "cli/app/cmd/init"
	"log"
	"github.com/spf13/cobra"

	"cli/app/cmd/pr"
)

var rootCmd = &cobra.Command{
	Use:   "bbc",
	Short: "BBC is a Bitbucket Cloud CLI",
	Long:  "A fast and flexible CLI for Bitbucket Cloud, so you don't have to deal with the shitty UI",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(pr.Cmd)
	rootCmd.AddCommand(initialize.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
