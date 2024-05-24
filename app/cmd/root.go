package cmd

import (
	"log"

	"github.com/spf13/cobra"

	initialize "cli/app/cmd/init"

	"cli/app/cmd/pr"
)

var rootCmd = &cobra.Command{
	Use:   "bbcli",
	Short: "BBC is a Bitbucket Cloud CLI",
	Long:  "A fast and flexible CLI for Bitbucket Cloud, so you don't have to deal with the shitty UI",
	Run: func(cmd *cobra.Command, args []string) {
		println("Welcome to BBCLI, please use the --help flag to see the available commands")
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
