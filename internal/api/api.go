package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"mid-passport-checker/internal/config"
	"mid-passport-checker/internal/models"
	"net/http"
)

type IDataSource interface {
	GetDataFromServer(ids []string, ctx context.Context) ([]models.Passport, error)
}

type API struct {
	cfg config.Config
}

func NewAPI(cfg config.Config, logger zerolog.Logger) *API {
	log.Logger = logger
	return &API{cfg: cfg}
}

func (api *API) GetDataFromServer(ids []string, ctx context.Context) ([]models.Passport, error) {
	var passports []models.Passport
	for _, id := range ids {
		select {
		case <-ctx.Done():
			return nil, models.ErrExceededTimout
		default:
			client := &http.Client{}
			url := api.cfg.APIAddress + api.cfg.ReqAddress + id
			req, err := http.NewRequest(http.MethodGet, url, nil)
			req.Header.Set("User-Agent", "FooBar")
			req.Header.Set("Accept", "*/*")
			req.Header.Set("Connection", "keep-alive")

			resp, err := client.Do(req)
			if err != nil {
				return nil, fmt.Errorf("request: %w", err)
			}

			jsonBytes, err := io.ReadAll(resp.Body)
			if err != nil && !errors.Is(err, io.EOF) {
				return nil, fmt.Errorf("read response body: %w", err)
			}
			defer resp.Body.Close()

			var passport models.Passport
			if len(jsonBytes) != 0 {
				err = json.Unmarshal(jsonBytes, &passport)
				if err != nil {
					return nil, fmt.Errorf("json unmarshall: %w", err)
				}
			}
			passports = append(passports, passport)
		}
	}
	return passports, nil
}
