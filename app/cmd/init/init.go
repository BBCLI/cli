package init

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	config2 "cli/app/lib/config"
)

var Cmd = &cobra.Command{
	Use:   "init",
	Short: "Configure your access",
	Run: func(cmd *cobra.Command, args []string) {

		config, err := config2.GetConfig()
		if err != nil {
			return
		}

		username, err := getLine("username")
		if err != nil {
			return
		}
		pass, err := getLine("app password")
		if err != nil {
			return
		}

		config.Authorization.Username = *username
		config.Authorization.Password = *pass

		err = config2.SaveConfig(config)

		if err != nil {
			fmt.Println("Error saving the config!", err)
		}
	},
}

func getLine(name string) (*string, error) {
	fmt.Print("Please enter your " + name + ": ")
	var ret string
	_, err := fmt.Scanln(&ret)
	if err != nil {
		fmt.Println("Error Reading your " + name)
		return nil, err
	}
	log.Println("Your "+name+": ", ret)
	return &ret, nil
}
