package service

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"mid-passport-checker/internal/api"
	"mid-passport-checker/internal/config"
	"mid-passport-checker/internal/models"
	"mid-passport-checker/internal/storage"
)

type IService interface {
	PassportService() error
	GetPassports(ctx context.Context) ([]models.Passport, error)
	SavePassportInfo(passports []models.Passport, ctx context.Context) error
}

type Service struct {
	cfg config.Config
	str storage.IStorage
	api api.IDataSource
}

func NewService(cfg config.Config, str storage.IStorage, api api.IDataSource, logger zerolog.Logger) *Service {
	log.Logger = logger
	return &Service{
		cfg: cfg,
		str: str,
		api: api,
	}
}
