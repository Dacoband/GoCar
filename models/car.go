package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	CarID     uuid.UUID `json: "id"`
	Name      string    `json: "name"`
	Year      string    `json: "year"`
	Brand     string    `json: "brand"`
	FuelType  string    `json: "fuel_type"`
	Engine    Engine    `json: "engine"`
	Price     float64   `json: "price"`
	CreatedAt time.Time `json"created_at"`
	UpdatedAt time.Time `json"updated_at"`
}

type CarRequest struct {
	Name     string  `json: "name"`
	Year     string  `json: "year"`
	Brand    string  `json: "brand"`
	FuelType string  `json: "fuel_type"`
	Engine   Engine  `json: "engine"`
	Price    float64 `json: "price"`
}

func validateRequest(carRequest CarRequest) error {
	if err := validateName(carRequest.Name); err != nil {
		return err
	}
	if err := validateYear(carRequest.Year); err != nil {
		return err
	}
	if err := validateBrand(carRequest.Brand); err != nil {
		return err
	}
	if err := validateFuelType(carRequest.FuelType); err != nil {
		return err
	}
	if err := validateEngine(carRequest.Engine); err != nil {
		return err
	}
	if err := validatePrice(carRequest.Price); err != nil {
		return err
	}
	return nil
}

func validateName(name string) error {
	if name == "" {
		return errors.New("Name is required")
	}
	return nil
}
func validateYear(year string) error {
	if year == "" {
		return errors.New("Year is required")
	}
	_, err := strconv.Atoi(year)
	if err != nil {
		return errors.New("Year must be a valid number")
	}
	currentYear := time.Now().Year()
	yearInt, _ := strconv.Atoi(year)
	if yearInt < 1779 || yearInt > currentYear {
		return errors.New("Year must be between 1779 and current year")
	}
	return nil
}

func validateBrand(brand string) error {
	if brand == "" {
		return errors.New("Brand is required")
	}
	return nil
}

func validateFuelType(fuelType string) error {
	validateFuelTypes := []string{"Petrol", "Diesel", "Electric", "Hybird"}
	for _, validType := range validateFuelTypes {
		if fuelType == validType {
			return nil
		}
	}
	return errors.New("FuelType must be one of: Petrol, Diesel, Electric, Hybird")
}

func validateEngine(engine Engine) error {
	if engine.EngineID == uuid.Nil {
		return errors.New("EngineID is required")
	}
	if engine.Displacement <= 0 {
		return errors.New("Displacement must be greater than 0")
	}
	if engine.NoOfCyclinders <= 0 {
		return errors.New("Number of cyclinder must be greater than 0")
	}
	if engine.CarRange <= 0 {
		return errors.New("Car range must be greater than 0")
	}
	return nil
}

func validatePrice(price float64) error {
	if price <= 0 {
		return errors.New("Price must be greater than 0")
	}
	return nil
}
