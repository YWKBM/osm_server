package utils

import (
	"encoding/hex"
	"fmt"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/geojson"
)

func ImportGeoFromHex(hexGeo string) (string, error) {
	geoString, err := hex.DecodeString(hexGeo)
	if err != nil {
		return "", err
	}

	geo, err := wkb.Unmarshal(geoString)
	if err != nil {
		return "", err
	}

	polygon, ok := geo.(orb.Polygon)
	if !ok {
		return "", fmt.Errorf("не является полигоном")
	}

	return toGeoJSON(polygon), nil
}

func toGeoJSON(p orb.Polygon) string {
	f := geojson.NewFeature(p)
	data, _ := f.MarshalJSON()
	return string(data)
}
