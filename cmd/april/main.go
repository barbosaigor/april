package main

import (
	"github.com/barbosaigor/april/cli"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal("Error: %v\n", err)
	}
}
