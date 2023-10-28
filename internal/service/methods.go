package service

import (
	"context"
	"github.com/rs/zerolog/log"
	"mid-passport-checker/internal/models"
	"time"
)

func (svc *Service) PassportService() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	log.Info().Msg("Trying to get data")
	passports, err := svc.GetPassports(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get passports")
		return
	}

	log.Info().Msg("Data received. Trying to write.")
	err = svc.SavePassportInfo(passports)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to save passports")
	}
	return
}

func (svc *Service) GetPassports(ctx context.Context) ([]models.Passport, error) {
	ids, err := svc.str.GetPassportID(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get ids")
		return nil, err
	}
	passports, err := svc.api.GetDataFromServer(ids, ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get from server")
	}
	return passports, err
}

func (svc *Service) SavePassportInfo(passports []models.Passport) error {
	err := svc.str.PutPassportInfo(passports)
	return err
}
