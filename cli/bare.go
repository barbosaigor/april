package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"gitlab.com/barbosaigor/april"
)

var filepath string
var number uint32

func init() {
	bareCmd.Flags().StringVarP(&filepath, "file", "f", "conf.yml", "Configuration file")
	bareCmd.Flags().Uint32VarP(&number, "number", "n", 0, "Number of nodes to return")
	bareCmd.MarkFlagRequired("number")
	rootCmd.AddCommand(bareCmd)
}

var bareCmd = &cobra.Command{
	Use:   "bare",
	Short: "Bare execute only the internal picking algorithm",
	Long:  "Bare execute only the internal picking algorithm",
	Run: func(cmd *cobra.Command, args []string) {
		nodes, err := april.PickRandDepsYml(filepath, number)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(nodes)
	},
}
