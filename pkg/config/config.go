package config

import (
	"flag"
	"fmt"
	"os"
)

// Config has app configuration
type Config struct {
	Environment string
	AppHost     string
	AppPort     string
}

// NewConfig returns a pointer to a populated config object
func NewConfig() *Config {
	conf := &Config{}

	flag.StringVar(&conf.Environment, "environment", getEnv("APP_ENVIRONMENT", "development"), "app development")
	flag.StringVar(&conf.AppHost, "apphost", getEnv("DASHBOARDSERVER_HOST", "0.0.0.0"), "app host")
	flag.StringVar(&conf.AppPort, "appport", getEnv("DASHBOARDSERVER_PORT", "8000"), "app port")

	flag.Parse()

	return conf
}

// APIAddr Returns the Address of the service
func (c *Config) APIAddr() string {
	return fmt.Sprintf(":%s", c.AppPort)
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}
