package handler

import (
	"osm_server/config"
	"osm_server/features"
	"osm_server/handler/zone"

	"github.com/gorilla/mux"
)

type Handler struct {
	config      *config.Config
	zoneHandler *zone.ZoneHandler
}

func NewHandler(features features.Features, config *config.Config) *Handler {
	return &Handler{
		config:      config,
		zoneHandler: zone.NewZoneHandler(features.Zone),
	}
}

func (h *Handler) Init() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/zones/create", h.zoneHandler.CreateZone).Methods("POST")

	return router
}
