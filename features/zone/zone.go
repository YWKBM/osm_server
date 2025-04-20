package zone

import (
	"fmt"
	"osm_server/entities"
	"osm_server/repo/zone"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

type ZoneFeatures struct {
	repo *zone.ZoneRepo
}

func NewZoneFeatures(zoneRepo *zone.ZoneRepo) *ZoneFeatures {
	return &ZoneFeatures{
		repo: zoneRepo,
	}
}

func (z *ZoneFeatures) CreateZone(name, geoJson string) (int, error) {
	feature, err := geojson.UnmarshalFeature([]byte(geoJson))
	if err != nil {
		return 0, fmt.Errorf("unable to parse geoJson: %w", err)
	}

	polygon, ok := feature.Geometry.(orb.Polygon)
	if !ok {
		return 0, fmt.Errorf("is not polygon")
	}

	zone := entities.Zone{
		Name: name,
		Geo:  polygon,
	}

	id, err := z.repo.CreateZone(zone)
	if err != nil {
		return 0, err
	}

	return id, nil
}
