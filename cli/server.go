package cli

import (
	"fmt"

	"github.com/barbosaigor/april/server"
	"github.com/spf13/cobra"
)

var dstrHost string

func init() {
	serverCmd.Flags().IntVarP(&port, "port", "p", 7070, "Server port")
	serverCmd.Flags().StringVarP(&dstrHost, "chaos", "c", "", "Chaos server url")
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "April's Create a Server for API access",
	Long:  "April's Create a API Server. Listening on port 7111",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("(HTTP) Listening on port: ", port)
		if dstrHost != "" {
			server.SetChaosServerHost(dstrHost)
		}
		server.Serve(port)
	},
}
