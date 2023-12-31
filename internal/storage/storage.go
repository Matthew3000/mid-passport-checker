package storage

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"mid-passport-checker/internal/config"
	"mid-passport-checker/internal/models"
	"sync"
)

type IStorage interface {
	GetPassportID(ctx context.Context) ([]string, error)
	PutPassportInfo(passports []models.Passport) error
}

type Storage struct {
	cfg   config.Config
	mutex sync.Mutex
}

func NewStorage(cfg config.Config, logger zerolog.Logger) *Storage {
	log.Logger = logger
	return &Storage{
		cfg: cfg,
	}
}
