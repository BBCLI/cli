package init

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"cli/app/lib/auth"
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Enter your bitbucket cli token",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Please enter your token:")
		token := ""
		_, err := fmt.Scanln(&token)
		if err != nil {
			fmt.Print("Error Reading your token")
			return
		}
		log.Print("Your Token: ", token)
		err = auth.SetToken(token)
		if err != nil {
			fmt.Println("Error Setting your Token!", err)
		}
	},
}
