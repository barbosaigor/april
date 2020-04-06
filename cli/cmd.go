package cli

import (
	"fmt"

	"github.com/barbosaigor/april"
	"github.com/barbosaigor/april/auth"
	"github.com/barbosaigor/april/destroyer/request"
	"github.com/spf13/cobra"
)

const VERSION = "0.1.1"

var filepath string
var number uint32
var host string
var port int
var username string
var password string

func init() {
	rootCmd.Flags().StringVarP(&filepath, "file", "f", "conf.yml", "Configuration file")
	rootCmd.Flags().Uint32VarP(&number, "number", "n", 0, "Maximum number of services to return")
	rootCmd.Flags().StringVarP(&host, "chaos", "c", "localhost:7071", "Chaos server url")
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	rootCmd.Flags().StringVarP(&password, "password", "s", "", "Password")
	rootCmd.MarkFlagRequired("number")
	rootCmd.MarkFlagRequired("username")
	rootCmd.MarkFlagRequired("password")
}

var rootCmd = &cobra.Command{
	Use:   "april",
	Short: "April is a chaos testing tool",
	Long:  "A fast and flexible tool for chaos testing.",
	Run: func(cmd *cobra.Command, args []string) {
		services, err := april.PickRandDepsYml(filepath, number)
		if err != nil {
			fmt.Println(err)
			return
		}
		token := auth.EncryptUser(username, password)
		err = request.ReqToDestroy(host, services, token)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Services destroyed: ", services)
	},
	Version: VERSION,
}

func Execute() error {
	return rootCmd.Execute()
}
