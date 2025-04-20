package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HOST string
	PORT string

	DB_CONFIG PostgresConfig
}

type PostgresConfig struct {
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
	SSL_MODE    string
}

func Init() *Config {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	host, _ := os.LookupEnv("HOST")
	port, _ := os.LookupEnv("PORT")

	dbUser, _ := os.LookupEnv("DB_USER")
	dbName, _ := os.LookupEnv("DB_NAME")
	sslMode, _ := os.LookupEnv("SSL_MODE")
	dbPort, _ := os.LookupEnv("DB_PORT")
	dbPass, _ := os.LookupEnv("DB_PASS")
	dbHost, _ := os.LookupEnv("DB_HOST")

	return &Config{
		HOST: host,
		PORT: port,

		DB_CONFIG: PostgresConfig{
			DB_USER:     dbUser,
			DB_NAME:     dbName,
			SSL_MODE:    sslMode,
			DB_PORT:     dbPort,
			DB_PASSWORD: dbPass,
			DB_HOST:     dbHost,
		},
	}

}
