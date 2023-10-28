package storage

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/rs/zerolog/log"
	"mid-passport-checker/internal/models"
	"os"
	"strconv"
)

func (str *Storage) PutPassportInfo(passports []models.Passport) error {
	str.mutex.Lock()
	defer str.mutex.Unlock()

	file, err := os.Create(str.cfg.OutputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"ID", "Reception Date", "Status", "Percentage"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, passport := range passports {
		record := []string{
			passport.Uid,
			passport.ReceptionDate,
			passport.InternalStatus.Name,
			strconv.Itoa(passport.InternalStatus.Percent),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	log.Info().Msg(fmt.Sprintf("succesfully done writing to %s", str.cfg.OutputFile))
	return nil
}

func (str *Storage) GetPassportID(ctx context.Context) ([]string, error) {
	file, err := os.Open(str.cfg.InputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var passportIDs []string
	for _, record := range records {
		select {
		case <-ctx.Done():
			return nil, models.ErrExceededTimout
		default:
			for _, field := range record {
				passportIDs = append(passportIDs, field)
			}
		}
	}
	return passportIDs, nil
}
