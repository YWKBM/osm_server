package zone

import (
	"encoding/hex"
	"fmt"
	"osm_server/dto"
	"osm_server/entities"
	"osm_server/repo/zone"
	"osm_server/utils"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
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

func (z *ZoneFeatures) Create(name, geoJson string) (int, error) {
	feature, err := geojson.UnmarshalFeature([]byte(geoJson))
	if err != nil {
		return 0, fmt.Errorf("unable to parse geoJson: %w", err)
	}

	polygon, ok := feature.Geometry.(orb.Polygon)
	if !ok {
		return 0, fmt.Errorf("is not polygon")
	}

	geo, err := wkb.Marshal(polygon)
	if err != nil {
		return 0, fmt.Errorf("geo parsing failed: %w", err)
	}

	hexWKB := hex.EncodeToString(geo)

	zone := entities.Zone{
		Name: name,
		Geo:  hexWKB,
	}

	id, err := z.repo.Create(zone)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (z *ZoneFeatures) Get(id int) (dto.Zone, error) {
	var zoneDto dto.Zone

	zone, err := z.repo.Get(id)
	if err != nil {
		return zoneDto, err
	}

	geo, err := utils.ImportGeoFromHex(zone.Geo)
	if err != nil {
		return zoneDto, err
	}

	zoneDto.Name = zone.Name
	zoneDto.Geo = geo

	return zoneDto, nil
}

func (z *ZoneFeatures) GetList(page, limit int) (dto.ZoneList, error) {
	var zoneList dto.ZoneList

	list, err := z.repo.GetList(page, limit)
	if err != nil {
		return zoneList, err
	}

	for _, l := range list {
		zoneItem := dto.ZoneListItem{
			Id:   l.Id,
			Name: l.Name,
		}

		zoneList.Items = append(zoneList.Items, zoneItem)
	}

	return zoneList, nil
}
