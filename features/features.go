package features

import (
	"osm_server/dto"
	"osm_server/features/zone"
	"osm_server/repo"
)

// type ProviderFeatures interface {
// 	CreateProvider(name string, geo orb.Polygon) int
// 	UpdateProvider(geo orb.Polygon) (int, error)
// 	GetProvider(providerId int) dto.Provider
// 	GetProviders(geo orb.Polygon) dto.ProvidersByZoneList
// }

type ZoneFeatures interface {
	Create(name, geoJson string) (int, error)
	Get(id int) (dto.Zone, error)
	GetList(page, limit int) (dto.ZoneList, error)
}

type Features struct {
	Zone ZoneFeatures
}

func NewFeatures(repo repo.Repo) Features {
	return Features{
		Zone: zone.NewZoneFeatures(repo.Zone),
	}
}
