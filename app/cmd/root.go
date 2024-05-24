package cmd

import (
	"cli/app/cmd/pr"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "bbc",
	Short: "BBC is a bitbucket cloud cli",
	Long:  `A Fast and Flexible CLI for bitbucket cloud so you dont have to deal with the shitty UI`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	rootCmd.AddCommand(pr.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
