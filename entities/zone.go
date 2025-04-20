package entities

import "github.com/paulmach/orb"

type Zone struct {
	Id        int
	Name      string
	Providers []Provider
	Geo       orb.Polygon
}
