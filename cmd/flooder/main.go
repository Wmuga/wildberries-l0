package main

import (
	"log"

	"github.com/wmuga/wildberries-l0/config"
	"github.com/wmuga/wildberries-l0/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	app.Flooder(cfg)
}
