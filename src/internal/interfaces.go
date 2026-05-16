package internal

import (
	"github.com/CesarSotnas/simple-crud-chi/src/internal/models"
)

type ChampionsRotationServiceInterface interface {
	GetChampionsRotations() (models.ChampionsRotationResponse, error)
}

type ChampionsRotationRepositoryInterface interface {
	GetChampionsRotations() (models.ChampionsRotationResponse, error)
}
