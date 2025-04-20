package zone

import (
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"osm_server/entities"

	"github.com/paulmach/orb/encoding/wkb"
)

type ZoneRepo struct {
	db *sql.DB
}

func NewZoneRepo(db *sql.DB) *ZoneRepo {
	return &ZoneRepo{
		db: db,
	}
}

func (z *ZoneRepo) CreateZone(zone entities.Zone) (int, error) {
	tx, err := z.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	geo, err := wkb.Marshal(zone.Geo)
	if err != nil {
		return 0, fmt.Errorf("geo parsing failed: %w", err)
	}

	hexWKB := hex.EncodeToString(geo)

	var intersects bool
	err = tx.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM zones
			WHERE ST_Intersects(geom, ST_SetSRID($1::geometry, 4326))
			)`, hexWKB).Scan(&intersects)

	if err != nil {
		return 0, fmt.Errorf("intersection check failed: %w", err)
	}

	if intersects {
		return 0, errors.New("zone intersects with existing zones")
	}

	var zoneId int
	err = tx.QueryRow(`
		INSERT INTO zones (name, geom)
		VALUES ($1 , ST_SetSRID($2::geometry, 4326))
		RETURNING id
	`, zone.Name, hexWKB).Scan(&zoneId)
	if err != nil {
		return 0, fmt.Errorf("failed to insert provider: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return zoneId, nil
}
