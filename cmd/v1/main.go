package main

import (
	"flag"
	"log"
	"pocket/shared"
	"pocket/shared/config"
)

func main() {
	config_filename := flag.String("config", "", "Relative or absolute path to config file.")
	flag.Parse()

	cfg := config.LoadConfig(*config_filename)

	pocketNode, err := shared.Create(cfg)
	if err != nil {
		log.Fatalf("Failed to create pocket node: %s", err)
	}

	if err = pocketNode.Start(); err != nil {
		log.Fatalf("Failed to start pocket node: %s", err)
	}
}
