package main

import (
	"sensei/api"
	"sensei/conf"
	"sensei/db"
	"sensei/svc"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Starting application")

	err := conf.Setup()
	if err != nil {
		log.Fatal().Err(err).Msg("Fatal error in configuration setup")
		return
	}
	err = db.Setup()
	if err != nil {
		log.Fatal().Err(err).Msg("Fatal error in db setup")
		return
	}
	svc.Setup()
	api.Start()
}
