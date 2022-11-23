package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config defines the application's settings
type Config struct {
	App     AppConfig      `mapstructure:"app" json:"app"`
	DB      DatabaseConfig `mapstructure:"db" json:"db"`
	Storage StorageConfig  `mapstructure:"storage" json:"storage"`
	Worker  WorkerSettings `mapstructure:"worker" json:"worker"`
}

type AppConfig struct {
	Port     string `mapstructure:"port" json:"port" default:"8080"`
	Name     string `mapstructure:"name" default:"GoStream" json:"app"`
	Version  string `mapstructure:"version" default:"v1" json:"version"`
	Mode     string `mapstructure:"mode" default:"dev" json:"mode"` // Mode should be dev or production
	Url      string `mapstructure:"url" default:"http:localhost:8080" json:"url"`
	Timezone string `mapstructure:"timezone" default:"UTC" json:"timezone"` //Application Timezone
}

// Relation Database Config [mysql]
type DatabaseConfig struct {
	Dialect  string `mapstructure:"-" json:"-" default:"mysql"`
	Host     string `mapstructure:"host" json:"host" default:"localhost" env:"Database_HOST"`
	Port     string `mapstructure:"port" json:"port" default:"3306" env:"Database_PORT"`
	User     string `mapstructure:"user" json:"user" default:"root" env:"Database_USER"`
	Password string `mapstructure:"password" json:"password" default:"" env:"Database_PASSWORD"`
	Database string `mapstructure:"database" json:"database" default:"gostream" env:"Database_NAME"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host" default:"localhost"`
	Port     int    `mapstructure:"port" json:"port" default:"5672"`
	Username string `mapstructure:"username" json:"username" default:"guest"`
	Password string `mapstructure:"password" json:"password" default:"guest"`
	Database string `mapstructure:"database" json:"database" default:"0"`
}

type CacheConfig struct {
	Host     string `mapstructure:"host" json:"host" default:"localhost"`
	Port     int    `mapstructure:"port" json:"port" default:"5672"`
	Username string `mapstructure:"username" json:"username" default:"guest"`
	Password string `mapstructure:"password" json:"password" default:"guest"`
	Database string `mapstructure:"database" json:"database" default:"0"`
}

type WorkerSettings struct {
	Concurrency uint `mapstructure:"concurrency" json:"concurrency" default:"10" env:"WORKER_CONCURRENCY"`
}

func NewConfig() *Config {

	cfg, err := loadConfig(".")

	if err != nil {
		log.Fatal("can't be load config.")
	}

	return cfg
}

func loadConfig(path string) (*Config, error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}
	config := Config{}
	err = viper.Unmarshal(&config)

	if err != nil {
		return nil, err
	}

	return &config, nil
}
