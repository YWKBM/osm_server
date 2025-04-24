package zone

import (
	"encoding/json"
	"net/http"
	"osm_server/features"
	"osm_server/handler/zone/dto"
)

type ZoneHandler struct {
	ZoneFeatures features.ZoneFeatures
}

func NewZoneHandler(zoneFeatures features.ZoneFeatures) *ZoneHandler {
	return &ZoneHandler{ZoneFeatures: zoneFeatures}
}

func (z *ZoneHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateZoneRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	geoString := string(req.Geom)

	id, err := z.ZoneFeatures.Create(req.Name, geoString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	result := dto.CreateZoneResponse{
		Id: id,
	}

	resp, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}

func (z *ZoneHandler) Get(w http.ResponseWriter, r *http.Request) {
	var req dto.GetZoneRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result, err := z.ZoneFeatures.Get(req.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	resp, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}

func (z *ZoneHandler) GetList(w http.ResponseWriter, r *http.Request) {
	var req dto.GetListRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	result, err := z.ZoneFeatures.GetList(req.Page, req.Limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	resp, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(resp)
}
