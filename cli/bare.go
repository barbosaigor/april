package cli

import (
	"github.com/barbosaigor/april"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	bareCmd.Flags().StringVarP(&filepath, "file", "f", "conf.yml", "Configuration file")
	bareCmd.Flags().Uint32VarP(&number, "number", "n", 0, "Maximum number of services to return")
	bareCmd.MarkFlagRequired("number")
	rootCmd.AddCommand(bareCmd)
}

var bareCmd = &cobra.Command{
	Use:   "bare",
	Short: "Bare execute only the internal picking algorithm",
	Long:  "Bare execute only the internal picking algorithm",
	Run: func(cmd *cobra.Command, args []string) {
		services, err := april.PickFromYaml(filepath, number)
		if err != nil {
			logrus.Error(err)
			return
		}
		logrus.Infof("Selected Services: %v", services)
	},
}
