package config

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	Http        HttpCfg
	Logger      LoggerSettings
	DB          DBSettings
	Cache       Redis
	SwaggerUI   SwaggerUIConfig
	CatalogUrls struct {
		Catalog string
	}
}

type HttpCfg struct {
	Port int `json:"port"`
	Gin  struct {
		ReleaseMode bool `json:"ReleaseMode"`
		UseLogger   bool `json:"UseLogger"`
		UseRecovery bool `json:"UseRecovery"`
	}
	ProfilingEnabled bool `json:"ProfilingEnabled"`
	StopTimeout      int  `json:"StopTimeout"`
}

type LoggerSettings struct {
	Component string
	MinLevel  string
	Writer    struct {
		Brokers []string
		Topic   string
	}
}

// DBSettings Конфигурация БД
type DBSettings struct {
	ConnectionString string
	LogMode          bool
	MaxOpenConns     int
	MaxIdleConns     int
}

// Redis Настройка кэша приложения
type Redis struct {
	Host           string
	Port           int
	DatabaseNumber int
	Password       string
	Keys           struct {
		Halykid RedisKey
	}
}

// RedisKey model
type RedisKey struct {
	Key        string        `json:"key"`
	Expiration time.Duration `json:"expiration"`
}

type UrlConfig struct {
	Url         string
	HttpTimeout time.Duration
}

// GetHostPort Возвращает <Host>:<Port>
func (r Redis) GetHostPort() string {
	return fmt.Sprintf("%v:%v", r.Host, r.Port)
}

// SwaggerUIConfig настроки swagger
type SwaggerUIConfig struct {
	PageTitle   string
	Host        string
	Description string
}

func Init(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(file)
	cfg := new(Config)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
