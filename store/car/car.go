package car

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Dacoband/GoCar/models"
	"github.com/google/uuid"
)

type Store struct {
	db *sql.DB
}

func new(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) GetCarById(ctx context.Context, id string) (models.Car, error) {
	var car models.Car
	query := `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at,
			  e.id, e.displacement, e.no_of_cyclinders, e.car_range FROM car c LEFT JOIN engine e ON c.engine_id = e.id WHERE c.id=$1`
	row := s.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&car.CarID,
		&car.Name,
		&car.Year,
		&car.Brand,
		&car.FuelType,
		&car.Engine.EngineID,
		&car.Price,
		&car.CreatedAt,
		&car.UpdatedAt,
		&car.Engine.EngineID,
		&car.Engine.Displacement,
		&car.Engine.NoOfCyclinders,
		&car.Engine.CarRange,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return car, nil
		}
		return car, err
	}
	return car, nil
}

func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	var cars []models.Car
	var query string
	if isEngine {
		query = `SELECT c.id, c.name, c.year, c.brand, c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at,
			  e.id, e.displacement, e.no_of_cyclinders, e.car_range FROM car c LEFT JOIN engine e ON c.engine_id = e.id WHERE c.brand =$1`
	} else {
		query = `SELECT id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at FROM car WHERE brand =$1`
	}
	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var car models.Car
		if isEngine {
			var engine models.Engine
			err := rows.Scan(
				&car.CarID,
				&car.Name,
				&car.Year,
				&car.Brand,
				&car.FuelType,
				&car.Engine.EngineID,
				&car.Price,
				&car.CreatedAt,
				&car.UpdatedAt,
				&car.Engine.EngineID,
				&car.Engine.Displacement,
				&car.Engine.NoOfCyclinders,
				&car.Engine.CarRange,
			)
			if err != nil {
				return nil, err
			}
			car.Engine = engine
		} else {
			err := rows.Scan(
				&car.CarID,
				&car.Name,
				&car.Year,
				&car.Brand,
				&car.FuelType,
				&car.Engine.EngineID,
				&car.Price,
				&car.CreatedAt,
				&car.UpdatedAt,
			)
			if err != nil {
				return nil, err
			}
		}
		cars = append(cars, car)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}
func (s Store) CreateCar(ctx context.Context, carRequest *models.CarRequest) (models.Car, error) {
	var createdCar models.Car
	var engineID uuid.UUID
	err := s.db.QueryRowContext(ctx, "SELECT id FROM engine WHERE id=$1", carRequest.Engine.EngineID).Scan(&engineID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return createdCar, errors.New("engine_id does not exist in the engine table")
		}
		return createdCar, err
	}
	carID := uuid.New()
	createdAt := time.Now()
	updatedAt := createdAt

	newCar := models.Car{
		CarID:     carID,
		Name:      carRequest.Name,
		Year:      carRequest.Year,
		Brand:     carRequest.Brand,
		FuelType:  carRequest.FuelType,
		Engine:    carRequest.Engine,
		Price:     carRequest.Price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	// transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return createdCar, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()
	query := `INSERT INTO car(id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at)
			 VALUES ($1, $2, &3, $4, $5, $6, $7, $8, $9)
			 RETURNING id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at`
	err = tx.QueryRowContext(ctx, query,
		newCar.CarID,
		newCar.Name,
		newCar.Year,
		newCar.Brand,
		newCar.FuelType,
		newCar.Engine.EngineID,
		newCar.Price,
		newCar.CreatedAt,
		newCar.UpdatedAt,
	).Scan(
		&createdCar.CarID,
		&createdCar.Name,
		&createdCar.Year,
		&createdCar.Brand,
		&createdCar.FuelType,
		&createdCar.Engine.EngineID,
		&createdCar.Price,
		&createdCar.CreatedAt,
		&createdCar.UpdatedAt,
	)
	if err != nil {
		return createdCar, nil
	}
	return createdCar, nil
}

func (s Store) UpdateCarById(ctx context.Context, id string, carRequest *models.CarRequest) (models.Car, error) {

}

func (s Store) DeleteCarById(ctx context.Context, id string) (models.Car, error) {

}
