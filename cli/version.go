package cli

import (
	"fmt"

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
		fmt.Printf("April %v\n", VERSION)
	},
}
