package me

import (
	"fmt"

	"github.com/spf13/cobra"

	"cli/app/lib/user"
)

var Cmd = &cobra.Command{
	Use:   "me",
	Short: "user",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := user.CurrentUser()
		if err != nil {
			return err
		}
		fmt.Println("UUID: ", *res.Uuid)
		fmt.Println("Display Name: ", *res.DisplayName)
		fmt.Println("Type: ", res.Type)
		fmt.Println("CreatedOn: ", res.CreatedOn)
		return nil
	},
}
