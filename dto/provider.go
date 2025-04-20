package dto

type Provider struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	ZoneName string `json:"zone_name"`
}

type ProvidersByZoneList struct {
	ZoneName  string     `json:"zone_name"`
	Providers []Provider `json:"providers"`
}
