package cli

import (
	"fmt"

	"github.com/barbosaigor/april/auth"
	"github.com/barbosaigor/april/internal/chaoshost"
	"github.com/spf13/cobra"
)

// VERSION describes current Aprils version
const VERSION = "0.8.0"

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
		token := auth.EncryptUser(username, password)
		ch := chaoshost.ChaosHost{Host: host, Token: token}
		svs, err := ch.PickAndShutdownInstancesFile(filepath, number)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Selected Services: ", svs)
	},
	Version: VERSION,
}

// Execute execute CLI operations
func Execute() error {
	return rootCmd.Execute()
}
