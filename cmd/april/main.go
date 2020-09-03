package main

import (
	"os"

	"github.com/barbosaigor/april/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
