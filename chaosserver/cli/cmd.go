package cli

import (
	cs "github.com/barbosaigor/april/chaosserver"
	"github.com/spf13/cobra"
)

// VERSION contain current version of chaos server
const VERSION = "1.0.0"

var username string
var password string
var port int

// Cs used as stretegy design pattern, allowing easily switch chaos server implementations
var Cs cs.ChaosServer

func init() {
	RootCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	RootCmd.Flags().StringVarP(&password, "password", "s", "", "Password")
	RootCmd.Flags().IntVarP(&port, "port", "p", 7071, "Server port number")
	RootCmd.MarkFlagRequired("username")
	RootCmd.MarkFlagRequired("password")
}

// RootCmd is an command structure used in CLI implementation
var RootCmd = &cobra.Command{
	Use:   "chaos",
	Short: "Chaos server terminates instances via an API.",
	Long:  "Chaos server terminates instances via an API.",
	Run: func(cmd *cobra.Command, args []string) {
		serv := cs.New(port, Cs)
		serv.Cred.Register(username, password)
		serv.Serve()
	},
	Version: VERSION,
}

// Execute runs CLI operations
func Execute() error {
	return RootCmd.Execute()
}
