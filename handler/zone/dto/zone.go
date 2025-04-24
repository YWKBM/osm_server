package dto

import "encoding/json"

type CreateZoneRequest struct {
	Name string          `json:"name"`
	Geom json.RawMessage `json:"geo"`
}

type CreateZoneResponse struct {
	Id int `json:"id"`
}

type GetZoneRequest struct {
	Id int `json:"Id"`
}

type GetListRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
