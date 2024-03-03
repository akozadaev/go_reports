package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"log"
	"path/filepath"
	"time"
)

type Config struct {
	ServerConfig  ServerConfig  `json:"server" yaml:"server"`
	DBConfig      DBConfig      `json:"db"`
	LoggingConfig LoggingConfig `json:"logging" yaml:"logging"`
}

type LoggingConfig struct {
	Level       int    `json:"level"`
	Encoding    string `json:"encoding"`
	Development bool   `json:"development"`
}

type LinksConfig struct {
	LifeTime time.Duration `json:"lifetime"`
}

type ServerConfig struct {
	Port             int           `json:"port"`
	ReadTimeout      time.Duration `json:"readTimeout"`
	WriteTimeout     time.Duration `json:"writeTimeout"`
	GracefulShutdown time.Duration `json:"gracefulShutdown"`
	Host             string        `json:"host"`
}

type DBConfig struct {
	DataSourceName string `json:"dataSourceName"`
	LogLevel       int    `json:"logLevel"`
	Pool           struct {
		MaxOpen     int           `json:"maxOpen"`
		MaxIdle     int           `json:"maxIdle"`
		MaxLifetime time.Duration `json:"maxLifetime"`
	} `json:"pool"`
}

func Load() (*Config, error) {
	k := koanf.New(".")

	path, err := filepath.Abs("./config/config.local.yaml")
	if err != nil {
		log.Printf("failed to get absoulute config path. configPath:%s, err: %v", "./config/config.local.yaml", err)
		return nil, err
	}

	log.Printf("load config file from %s", path)
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Printf("failed to load config from file. err: %v", err)
		return nil, err
	}

	var cfg Config
	if err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{Tag: "json", FlatPaths: false}); err != nil {
		log.Printf("failed to unmarshal with conf. err: %v", err)
		return nil, err
	}
	return &cfg, err
}