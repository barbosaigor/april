package cli

import (
	"github.com/barbosaigor/april/server"
	"github.com/sirupsen/logrus"
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
		if dstrHost != "" {
			server.SetChaosServerHost(dstrHost)
		}
		logrus.Infof("(HTTP) Listening on port: %d", port)
		server.Serve(port)
	},
}
