package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID       uuid.UUID `json: "id"`
	Displacement   int64     `json: "displacement"`
	NoOfCyclinders int64     `json: "noOfCyclinders"`
	CarRange       int64     `json: "carRange"`
}

type EngineRequest struct {
	Displacement   int64 `json: "displacement"`
	NoOfCyclinders int64 `json: "noOfCyclinders"`
	CarRange       int64 `json: "carRange"`
}

func validateEngineRequest(engineRequest EngineRequest) error {
	if err := validDisplacement(engineRequest.Displacement); err != nil {
		return err
	}
	if err := validateNoOfCyclinders(engineRequest.NoOfCyclinders); err != nil {
		return err
	}
	if err := validateCarRange(engineRequest.CarRange); err != nil {
		return err
	}
	return nil
}

func validDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("Displacement must be geater than 0")
	}
	return nil
}

func validateNoOfCyclinders(noOfCyclinders int64) error {
	if noOfCyclinders <= 0 {
		return errors.New("NoOfCyclinders must be geater than 0")
	}
	return nil
}

func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("CarRange must be geater than 0")
	}
	return nil
}
