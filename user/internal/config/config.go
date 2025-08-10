package config

import "time"

type Config struct {
	Service struct {
		Host         string        `yaml:"host"`
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"read_timeout"`
		WriteTimeout time.Duration `yaml:"write_timeout"`
		IdleTimeout  time.Duration `yaml:"idle_timeout"`
	} `yaml:"service"`
	DB struct {
		Host         string `yaml:"host"`
		Port         string `yaml:"port"`
		User         string `yaml:"user"`
		Password     string `yaml:"password"`
		DBName       string `yaml:"db_name"`
		Shema        string `yaml:"schema"`
		AppName      string `yaml:"app_name"`
		MaxIdleConns int    `yaml:"max_idle_conns"`
		MaxOpenConns int    `yaml:"max_open_conns"`
	} `yaml:"db"`
}
