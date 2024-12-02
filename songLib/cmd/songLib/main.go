package main

import (
	"log"

	"github.com/jMurad/musicService/songLib/internal/app"
	"github.com/jMurad/musicService/songLib/internal/config"
)

func main() {
	// Configuration
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatal(err)
	}

	app := app.New(cfg)

	// Start
	app.Start()

}
