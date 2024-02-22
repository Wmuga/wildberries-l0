package app

import (
	"log"

	"github.com/wmuga/wildberries-l0/config"
)

func L0(config *config.Config) {
	log.Println(*config)
}
