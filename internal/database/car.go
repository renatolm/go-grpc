package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Car struct {
	db             *sql.DB
	ID             string
	Name           string
	HorsePower     *int
	ManufacturerID string
}

func NewCar(db *sql.DB) *Car {
	return &Car{db: db}
}

func (c *Car) Create(name string, horse_power int, manufacturerID string) (*Car, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO cars (id, name, horse_power, manufacturer_id) VALUES ($1, $2, $3, $4)", id, name, horse_power, manufacturerID)
	if err != nil {
		return nil, err
	}
	return &Car{
		ID:             id,
		Name:           name,
		HorsePower:     &horse_power,
		ManufacturerID: manufacturerID,
	}, nil
}

func (c *Car) FindAll() ([]Car, error) {
	rows, err := c.db.Query("SELECT id, name, horse_power, manufacturer_id FROM cars")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cars := []Car{}
	for rows.Next() {
		var id, name, manufacturerID string
		var horse_power *int
		if err := rows.Scan(&id, &name, &horse_power, &manufacturerID); err != nil {
			return nil, err
		}
		cars = append(cars, Car{ID: id, Name: name, HorsePower: horse_power, ManufacturerID: manufacturerID})
	}
	return cars, nil
}

func (c *Car) FindByManufacturerID(manufacturerID string) ([]Car, error) {
	rows, err := c.db.Query("SELECT id, name, horse_power, manufacturer_id FROM cars WHERE manufacturer_id = $1", manufacturerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cars := []Car{}
	for rows.Next() {
		var id, name, manufacturerID string
		var horse_power *int
		if err := rows.Scan(&id, &name, &horse_power, &manufacturerID); err != nil {
			return nil, err
		}
		cars = append(cars, Car{ID: id, Name: name, HorsePower: horse_power, ManufacturerID: manufacturerID})
	}
	return cars, nil
}
