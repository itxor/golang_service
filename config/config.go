package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"

	"github.com/jinzhu/configor"
)

// Config ...
type Config struct {
	PgURL      string `env:"PG_URL"`
	PgProto    string `env:"PG_PROTO"`
	PgAddr     string `env:"PG_ADDR"`
	PgDb       string `env:"PG_DB"`
	PgUser     string `env:"PG_USER"`
	PgPassword string `env:"PG_PASSWORD"`
}

var (
	config Config
	once   sync.Once
)

// Get ...
func Get() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}

		envType := os.Getenv("ENV")
		if envType == "" {
			envType = "dev"
		}
		if err := configor.New(&configor.Config{
			Environment: envType,
		}).Load(
			&config,
			"config.json",
		); err != nil {
			log.Fatal(err)
		}
		configBytes, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Configurations:", string(configBytes))
	})

	return &config
}
