package main

import (
	"fmt"
	"log"
	"net/http"
	"osm_server/config"
	"osm_server/database"
	"osm_server/features"
	"osm_server/handler"
	"osm_server/repo"

	"github.com/pressly/goose"
)

func main() {
	config := config.Init()

	connectionInfo := database.ConnectionInfo{
		Host:     config.DB_CONFIG.DB_HOST,
		Port:     config.DB_CONFIG.DB_PORT,
		Username: config.DB_CONFIG.DB_USER,
		Password: config.DB_CONFIG.DB_PASSWORD,
		DBName:   config.DB_CONFIG.DB_NAME,
		SSLMode:  config.DB_CONFIG.SSL_MODE,
	}

	db, err := database.NewPostgresConnection(connectionInfo)
	if err != nil {
		panic(err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	defer db.Close()

	repo := repo.NewRepo(db)
	features := features.NewFeatures(*repo)
	handler := handler.NewHandler(features, config)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", config.HOST, config.PORT), handler.Init())

	if err != nil {
		log.Fatal(err)
	}
}
