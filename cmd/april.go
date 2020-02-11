package main

import (
	"fmt"

	"gitlab.com/barbosaigor/april/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
