package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"mid-passport-checker/internal/api"
	"mid-passport-checker/internal/config"
	"mid-passport-checker/internal/service"
	"mid-passport-checker/internal/storage"
	"mid-passport-checker/internal/tools"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func main() {
	cfg := config.InitiateConfig()
	fmt.Println("Configuration loaded")
	logger := tools.NewLogger()
	log.Info().Msg("successfully started logger")

	str := storage.NewStorage(cfg, logger.Logger)
	api := api.NewAPI(cfg, logger.Logger)
	svc := service.NewService(cfg, str, api, logger.Logger)

	svc.PassportService()
	timer, err := strconv.Atoi(cfg.ReqTimer)
	if err != nil {
		log.Fatal().Err(err).Msg("failed convert timer")
	}
	log.Info().Msg(fmt.Sprintf("timer is set to %s minutes", cfg.ReqTimer))

	tickerUpdate := time.NewTicker(time.Minute * time.Duration(timer))
	go func() {
		for range tickerUpdate.C {
			svc.PassportService()
		}
	}()
	log.Info().Msg("Press Ctrl+C to exit")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Received user interrupt. Exiting.")
}
