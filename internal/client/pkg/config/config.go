package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

// contextKey const
type contextKey string

const (
	CookieKey  = contextKey("cookie")
	CookiePath = contextKey("/")
)

func (c contextKey) String() string {
	return string(c)
}

// Config struct
type Config struct {
	Request       string `env:"REQUEST_HANDLER" envDefault:"login"`
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"http://127.0.0.1:9090"`
	Path          string `env:"PATH_JSON" envDefault:"load.json"`
	CookieKey     string `env:"COOKIE_KEY" envDefault:"xxx"`
}

// Parse method of Config
func (c *Config) Parse() error {
	flag.StringVar(&c.ServerAddress, "a", "", "Client address")
	flag.StringVar(&c.Request, "r", "", "Request handler")
	flag.StringVar(&c.Path, "p", "", "Path for JSON file")
	flag.StringVar(&c.Path, "k", "", "Cookie key for auth entry")
	flag.Parse()

	//settings redefinition, if env variables are used
	err := env.Parse(c)

	return err
}
