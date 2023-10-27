package tools

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type Logger struct {
	Logger  zerolog.Logger
	timer   *time.Timer
	logFile *os.File
	dir     string
}

func NewLogger() *Logger {
	logger, err := instantiateLogger()
	if err != nil {
		log.Fatal().Err(err).Msg("error instantiating the logger")
		return nil
	}
	return &Logger{
		Logger: logger,
	}
}

func instantiateLogger() (zerolog.Logger, error) {
	logWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	logger := zerolog.New(logWriter).With().Timestamp().Logger()

	zerolog.SetGlobalLevel(1)
	log.Logger = logger

	return logger, nil
}
