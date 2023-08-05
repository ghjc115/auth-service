package config

import (
	"os"
	"time"
)

type Config struct {
	Env        string `yaml:"env" env-default:"dev"`
	Version    string `yaml:"version" env-default:"0.0.1"`
	Database   `yaml:"dao"`
	HttpServer `yaml:"http_server"`
}

type Database struct {
	Host     string `yaml:"host" env-default:"mongodb://localhost"`
	Username string `yaml:"username" env-default:"root"`
	Password string `yaml:"password" env-default:"root"`
	Port     string `yaml:"port" env-default:"27017"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"0.0.0.0:7500"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() Config {
	env := os.Getenv("ENV")
	version := os.Getenv("VERSION")
	databaseHost := os.Getenv("DATABASE_HOST")
	databaseUsername := os.Getenv("DATABASE_USERNAME")
	databasePassword := os.Getenv("DATABASE_PASSWORD")
	databasePort := os.Getenv("DATABASE_PORT")
	httpAddress := os.Getenv("HTTP_ADDRESS")
	httpTimeout, _ := time.ParseDuration(os.Getenv("HTTP_TIMEOUT"))
	httpIdleTimeout, _ := time.ParseDuration(os.Getenv("HTTP_IDLE_TIMEOUT"))

	return Config{
		Env:     env,
		Version: version,
		Database: Database{
			Host:     databaseHost,
			Username: databaseUsername,
			Password: databasePassword,
			Port:     databasePort,
		},
		HttpServer: HttpServer{
			Address:     httpAddress,
			Timeout:     httpTimeout,
			IdleTimeout: httpIdleTimeout,
		},
	}
}
