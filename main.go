package main

import (
	"fmt"
	"log"

	internal "github.com/JaygnatBuilds/gator_aggregator/internal/config"
)

func main() {

	// initialize config
	config, err := internal.ReadConfig()
	if err != nil {
		log.Fatalf("Error initializing config : %v", err)
	}

	fmt.Println(config)

	config.SetUser("justin")

	// read config after username setting
	config, err = internal.ReadConfig()
	if err != nil {
		log.Fatalf("Error initializing config : %v", err)
	}

	fmt.Println(config)
}
