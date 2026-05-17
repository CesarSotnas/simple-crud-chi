package ports

import (
	"github.com/CesarSotnas/simple-crud-chi/internal/model"
)

type ChampionsRotationServiceInterface interface {
	GetChampionsRotations() (model.ChampionsRotationResponse, error)
}

type ChampionsRotationRepositoryInterface interface {
	GetChampionsRotations() (model.ChampionsRotationResponse, error)
}
