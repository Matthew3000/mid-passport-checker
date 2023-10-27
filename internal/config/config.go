package config

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"github.com/pelletier/go-toml"
	"github.com/rs/zerolog/log"
)

type Config struct {
	ConfigPath string `env:"CONFIG_PATH"     envDefault:"config.toml"`
	OutputFile string `env:"OUTPUT_FILE"     envDefault:"output/passports.csv"`
	InputFile  string `env:"OUTPUT_FILE"     envDefault:"input/ids.csv"`
	APIAddress string `env:"SERVER_ADDRESS"  envDefault:"https://info.midpass.ru/"`
	ReqAddress string `env:"REQUEST_ADDRESS" envDefault:"api/request/"`
	ReqTimer   string `env:"REQUEST_TIMER"   envDefault:"10"`
}

func InitiateConfig() Config {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		log.Error().Err(err).Msg("error parsing flags")
	}
	flag.StringVar(&cfg.ConfigPath, "c", cfg.ConfigPath, "Path to config file")
	flag.StringVar(&cfg.OutputFile, "o", cfg.OutputFile, "Output file")
	flag.StringVar(&cfg.InputFile, "i", cfg.InputFile, "Input file")
	flag.StringVar(&cfg.APIAddress, "a", cfg.APIAddress, "Server address")
	flag.StringVar(&cfg.ReqAddress, "r", cfg.ReqAddress, "Request url")
	flag.StringVar(&cfg.ReqTimer, "t", cfg.ReqTimer, "timer")
	flag.Parse()

	fileData, err := toml.LoadFile(cfg.ConfigPath)
	if err != nil {
		log.Error().Err(err).Msg("error parsing config file")
		return cfg
	}

	var tomlCfg Config
	if err = fileData.Unmarshal(&tomlCfg); err != nil {
		log.Error().Err(err).Msg("error unmarshalling config data")
		return cfg
	}

	if tomlCfg.OutputFile != "" {
		cfg.OutputFile = tomlCfg.OutputFile
	}
	if tomlCfg.InputFile != "" {
		cfg.InputFile = tomlCfg.InputFile
	}
	if tomlCfg.APIAddress != "" {
		cfg.APIAddress = tomlCfg.APIAddress
	}
	if tomlCfg.ReqAddress != "" {
		cfg.ReqAddress = tomlCfg.ReqAddress
	}
	if tomlCfg.ReqTimer != "" {
		cfg.ReqTimer = tomlCfg.ReqTimer
	}
	return cfg
}
