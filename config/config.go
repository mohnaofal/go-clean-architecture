package config

import (
	"os"
	"strconv"

	"github.com/mohnaofal/clean-architecture/pkg/database"
)

type Config struct {
	port int
	db   database.GormConnector
}

func InitConfig() *Config {
	cfg := new(Config)
	LoadEnv(cfg)
	cfg.db = database.InitGorm()

	return cfg
}

func LoadEnv(cfg *Config) {
	// PORT env
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic("Invalid PORT env")
	}
	cfg.port = port
}

func (c *Config) PORT() int {
	return c.port
}

func (c *Config) DB() database.GormConnector {
	return c.db
}
