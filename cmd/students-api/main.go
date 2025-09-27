package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Httpserver holds server-specific config
type Httpserver struct {
	Addr string `yaml:"addr" env:"HTTP_ADDR" `
}

// Config is the root application config
type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true"`
	Storagepath string     `yaml:"storage_path" env:"STORAGE_PATH" env-required:"true"`
	Httpserver  Httpserver `yaml:"http_server"`
}

// MustLoad loads config from ENV or file and exits on failure
func MustLoad() *Config {
	var configpath string

	// 1. Check environment variable
	configpath = os.Getenv("CONFIG_PATH")

	// 2. If not set, check CLI flag
	if configpath == "" {
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configpath = *flags
		if configpath == "" {
			log.Fatal("config path is not set (use CONFIG_PATH env var or --config flag)")
		}
	}

	// 3. Ensure the file exists
	if _, err := os.Stat(configpath); os.IsNotExist(err) {
		log.Fatalf("config file doesn't exist: %s", configpath)
	}

	// 4. Parse the config
	var cfg Config
	err := cleanenv.ReadConfig(configpath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file: %s", err.Error())
	}

	return &cfg
}
