package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

// contextKey const
type contextKey string

const (
	CookieKey = contextKey("cookie")
)

func (c contextKey) String() string {
	return string(c)
}

type Config struct {
	ServerAddress string `env:"RUN_ADDRESS" envDefault:":9090"`
	DatabaseDSN   string `env:"DATABASE_URI" envDefault:"postgres://dmosk:dmosk@localhost:5432/dmosk?sslmode=disable"`
	PasswordSalt  string `env:"PASSWORD_SALT" envDefault:"0000"`
}

func (c *Config) Parse() error {
	flag.StringVar(&c.ServerAddress, "a", "", "Server address")
	flag.StringVar(&c.DatabaseDSN, "d", "", "Database URL conn")
	flag.StringVar(&c.PasswordSalt, "p", "", "Password salt")
	flag.Parse()

	//settings redefinition, if env variables are used
	err := env.Parse(c)

	return err
}
