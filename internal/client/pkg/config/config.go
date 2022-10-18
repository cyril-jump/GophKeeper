package config

import (
	"flag"

	"github.com/caarlos0/env/v6"
)

// Config struct
type Config struct {
	Request       string `env:"REQUEST_HANDLER"`
	ClientAddress string `env:"RUN_CLIENT" envDefault:":7070"`
	Path          string `env:"PATH_JSON"`
	CookieKey     string `env:"COOKIE_KEY"`
}

// Parse method of Config
func (c *Config) Parse() error {
	flag.StringVar(&c.ClientAddress, "a", "", "Client address")
	flag.StringVar(&c.Request, "r", "", "Request handler")
	flag.StringVar(&c.Path, "p", "", "Path for JSON file")
	flag.StringVar(&c.Path, "k", "", "Cookie key for auth entry")
	flag.Parse()

	//settings redefinition, if env variables are used
	err := env.Parse(c)

	return err
}
