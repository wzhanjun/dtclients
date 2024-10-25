package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

var Cfg envconfig

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	if err := env.Parse(&Cfg); err != nil {
		log.Println(err)
	}

	log.Printf("%#v\n", Cfg)
}

type envconfig struct {
	Words struct {
		AppId   string `env:"GRPC_WORDS_APP_ID"`
		Address string `env:"GRPC_WORDS_ADDRESS"`
	}
	Captcha struct {
		AppId   string `env:"GRPC_CAPTCHA_APP_ID"`
		Address string `env:"GRPC_CAPTCHA_ADDRESS"`
	}
	Slog struct {
		AppId   string `env:"GRPC_SLOG_APP_ID"`
		EsIndex string `env:"GRPC_SLOG_ESINDEX"`
		Address string `env:"GRPC_SLOG_ADDRESS"`
	}
}
