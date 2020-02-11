package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/barbosaigor/april"
	"gitlab.com/barbosaigor/april/destroyer/request"
)

const VERSION = "a1.0.0"
var host string

func init() {
	rootCmd.Flags().StringVarP(&filepath, "file", "f", "conf.yml", "Configuration file")
	rootCmd.Flags().Uint32VarP(&number, "number", "n", 0, "Number of nodes to return")
	rootCmd.Flags().StringVarP(&host, "host", "u", "localhost:7071", "Chaos server url")
	rootCmd.MarkFlagRequired("number")
}

var rootCmd = &cobra.Command{
	Use:   "april",
	Short: "April is a chaos testing tool",
	Long: "A fast and flexible tool for chaos testing.",
	Run: func(cmd *cobra.Command, args []string) {
		nodes, err := april.PickRandDepsYml(filepath, number)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = request.ReqToDestroy(host, nodes)
		if err != nil {
			fmt.Println(err.Error())
			return	
		}
		fmt.Println("Nodes destroyed: ", nodes)
	},
	Version: VERSION,
}

func Execute() error {
	return rootCmd.Execute()
}
