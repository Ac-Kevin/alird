package config

import (
	"os"

	"github.com/go-ini/ini"
	errors "github.com/pkg/errors"
)

// Config 配置文件
type Config struct {
	Ini *ini.File
}

// NewConfig def.
func NewConfig() (*Config, error) {
	filename := os.Getenv("CONFIG_FILE")
	if filename == "" {
		filename = "/alird/config.ini"
	}
	var cfg *ini.File
	cfg, err := ini.Load(filename)
	if err != nil {
		return nil, errors.WithMessagef(err, "load %s ini file error", filename)
	}
	return &Config{cfg}, nil
}
