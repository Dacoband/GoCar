package engine

import (
	"context"
	"database/sql"

	"github.com/Dacoband/GoCar/models"
)

type EngineStore struct {
	db *sql.DB
}

func New(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (e EngineStore) GetEngineById(ctx context.Context, id string) (models.Engine, error) {

}

func (e EngineStore) CreateEngine(ctx context.Context, engineRequest *models.EngineRequest) (models.Engine, error) {

}

func (e EngineStore) UpdateEngineById(ctx context.Context, id string, engineRequest *models.EngineRequest) (models.Engine, error) {

}
func (e EngineStore) DeleteEngineById(ctx context.Context, id string) (models.Engine, error) {

}
