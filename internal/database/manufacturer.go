package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Manufacturer struct {
	db            *sql.DB
	ID            string
	Name          string
	OriginCountry *string
}

func NewManufacturer(db *sql.DB) *Manufacturer {
	return &Manufacturer{db: db}
}

func (m *Manufacturer) Create(name string, origin_country string) (Manufacturer, error) {
	id := uuid.New().String()
	_, err := m.db.Exec("INSERT INTO manufacturer (id, name , origin_country) VALUES ($1, $2, $3)", id, name, origin_country)

	if err != nil {
		return Manufacturer{}, err
	}

	return Manufacturer{ID: id, Name: name, OriginCountry: &origin_country}, nil
}

func (m *Manufacturer) FindAll() ([]Manufacturer, error) {
	rows, err := m.db.Query("SELECT id, name, origin_country FROM manufacturer")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	manufacturer := []Manufacturer{}
	for rows.Next() {
		var id, name string
		var origin_country *string
		if err := rows.Scan(&id, &name, &origin_country); err != nil {
			return nil, err
		}
		manufacturer = append(manufacturer, Manufacturer{ID: id, Name: name, OriginCountry: origin_country})
	}
	return manufacturer, nil
}

func (m *Manufacturer) FindByCarID(carID string) (Manufacturer, error) {
	var id, name string
	var origin_country *string
	err := m.db.QueryRow("SELECT m.id, m.name, m.origin_country FROM manufacturer m JOIN cars ca ON m.id = ca.manufacturer_id WHERE ca.id = $1", carID).Scan(&id, &name, &origin_country)
	if err != nil {
		return Manufacturer{}, err
	}
	return Manufacturer{ID: id, Name: name, OriginCountry: origin_country}, nil
}
