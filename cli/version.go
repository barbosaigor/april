package cli

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "April's version",
	Long:  "April is in alpha version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("April %v\n", VERSION)
	},
}
