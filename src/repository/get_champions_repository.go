package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/CesarSotnas/simple-crud-chi/src/helpers"
	"github.com/CesarSotnas/simple-crud-chi/src/internal"
	"github.com/CesarSotnas/simple-crud-chi/src/models"
)

type championsRotationsRepository struct{}

func NewGetChampionsRotationsRepository() internal.ChampionsRotationRepositoryInterface {
	return &championsRotationsRepository{}
}

func (r *championsRotationsRepository) GetChampionsRotations() (models.ChampionsRotationResponse, error) {
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://br1.api.riotgames.com/lol/platform/v3/champion-rotations?api_key=%s", apiKey)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	response, err := client.Get(url)
	if err != nil {
		errMsg := errors.New(helpers.GenericError)

		return models.ChampionsRotationResponse{}, errMsg
	}
	defer response.Body.Close()

	var data models.ChampionsRotationResponse

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		errMsg := errors.New(helpers.GenericError)

		return models.ChampionsRotationResponse{}, errMsg
	}

	return data, nil
}
