package repo

import (
	"database/sql"
	"osm_server/repo/zone"
)

type Repo struct {
	Zone *zone.ZoneRepo
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		Zone: zone.NewZoneRepo(db),
	}
}
