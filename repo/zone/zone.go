package zone

import (
	"database/sql"
	"errors"
	"fmt"
	"osm_server/entities"
)

type ZoneRepo struct {
	db *sql.DB
}

func NewZoneRepo(db *sql.DB) *ZoneRepo {
	return &ZoneRepo{
		db: db,
	}
}

func (z *ZoneRepo) Create(zone entities.Zone) (int, error) {
	tx, err := z.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	var intersects bool
	err = tx.QueryRow(`
		SELECT EXISTS(
			SELECT 1 FROM zones
			WHERE ST_Intersects(geom, ST_SetSRID($1::geometry, 4326))
			)`, zone.Geo).Scan(&intersects)

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
	`, zone.Name, zone.Geo).Scan(&zoneId)
	if err != nil {
		return 0, fmt.Errorf("failed to insert provider: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return zoneId, nil
}

func (z *ZoneRepo) Get(id int) (entities.Zone, error) {
	var zone entities.Zone

	err := z.db.QueryRow("SELECT * FROM zones WHERE Id = $1", id).Scan(&zone.Id, &zone.Name, &zone.Geo)
	if err != nil {
		return zone, err
	}

	return zone, nil
}

func (z *ZoneRepo) GetList(page, limit int) ([]entities.Zone, error) {
	rows, err := z.db.Query("SELECT * FROM zones LIMIT = $1 OFFSET $2", page, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	zone := entities.Zone{}
	zones := []entities.Zone{}

	for rows.Next() {
		err := rows.Scan(&zone.Id, &zone.Name, &zone.Geo)
		if err != nil {
			return nil, err
		}

		zones = append(zones, zone)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return zones, nil
}
