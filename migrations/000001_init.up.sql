-- +goose Up

CREATE EXTENSION IF NOT EXISTS postgis;


CREATE TABLE zones (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    geom GEOMETRY(Polygon, 4326) NOT NULL  -- WGS84 координаты
);

CREATE TABLE providers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    zone_id INTEGER NULL,  
    
    FOREIGN KEY (zone_id) 
        REFERENCES zones(id) 
        ON UPDATE CASCADE 
        ON DELETE SET NULL  
);
