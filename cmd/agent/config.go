package main

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port        string `env:"PORT" env-default:"8080"`
	Environment string `env:"ENV" env-default:"PRODUCTION"`
}

func NewConfig() (*Config, error) {
	var cfg Config

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}
