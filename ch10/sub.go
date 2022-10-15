package ch10

import (
	"fmt"

	"github.com/caarlos0/env"
)

func Sub() {
	list_9_4()
}

func list_9_4() {
	cfg, _ := New()
	fmt.Println(cfg)
}

type Config struct {
	Env  string `env:"TODO_ENV" envDefault:"dev"`
	Port int    `env:"PORT" envDefault:"80"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
