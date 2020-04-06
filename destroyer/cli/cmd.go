package cli

import (
	"github.com/barbosaigor/april/destroyer"
	"github.com/spf13/cobra"
)

const VERSION = "1.0.0"

var username string
var password string
var port int
var Cs destroyer.ChaosServer

func init() {
	RootCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	RootCmd.Flags().StringVarP(&password, "password", "s", "", "Password")
	RootCmd.Flags().IntVarP(&port, "port", "p", 7071, "Server port number")
	RootCmd.MarkFlagRequired("username")
	RootCmd.MarkFlagRequired("password")
}

var RootCmd = &cobra.Command{
	Use:   "chaos",
	Short: "Chaos server terminates instances via an API.",
	Long:  "Chaos server terminates instances via an API.",
	Run: func(cmd *cobra.Command, args []string) {
		serv := destroyer.New(port, Cs)
		serv.Cred.Register(username, password)
		serv.Serve()
	},
	Version: VERSION,
}

func Execute() error {
	return RootCmd.Execute()
}
