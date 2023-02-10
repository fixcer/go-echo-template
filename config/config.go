package config

import (
	"github.com/spf13/viper"
	"time"
)

type App struct {
	RuntimeRootPath string `mapstructure:"runtime-root-path"`
	Log             Log
}

type Log struct {
	SaveFile   bool   `mapstructure:"save-file"`
	SavePath   string `mapstructure:"save-path"`
	SaveName   string `mapstructure:"save-name"`
	FileExt    string `mapstructure:"file-ext"`
	TimeFormat string `mapstructure:"time-format"`
	MaxSize    int    `mapstructure:"max-size"`
	MaxBackup  int    `mapstructure:"max-backup"`
	MaxAge     int    `mapstructure:"max-age"`
}

type Server struct {
	RunMode      string `mapstructure:"run-mode"`
	Address      string
	ReadTimeout  time.Duration `mapstructure:"read-timeout"`
	WriteTimeout time.Duration `mapstructure:"write-timeout"`
}

type Database struct {
	Engine        string
	User          string
	Password      string
	Host          string
	Port          string
	Name          string
	MaxIdle       int    `mapstructure:"max-idle"`
	MaxOpen       int    `mapstructure:"max-open"`
	MigrationPath string `mapstructure:"migration-path"`
}

type Config struct {
	App      *App
	Server   *Server
	Database *Database
}

var Cfg *Config

func Setup() {
	viper.AutomaticEnv()
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		return
	}
}
