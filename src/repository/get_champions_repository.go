package repository

import (
	"fmt"
	"os"

	"github.com/CesarSotnas/simple-crud-chi/src/internal"
	"github.com/CesarSotnas/simple-crud-chi/src/internal/models"
)

type championsRotationsRepository struct{}

func NewGetChampionsRotationsRepository() internal.ChampionsRotationRepositoryInterface {
	return &championsRotationsRepository{}
}

func (r *championsRotationsRepository) GetChampionsRotations() (models.ChampionsRotationResponse, error) {
	apiKey := os.Getenv("API_KEY")
	_ = fmt.Sprintf("https://br1.api.riotgames.com/lol/platform/v3/champion-rotations?api_key=%s", apiKey)

	return models.ChampionsRotationResponse{}, nil
}
