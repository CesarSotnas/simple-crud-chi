package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/CesarSotnas/simple-crud-chi/internal/helpers"
	"github.com/CesarSotnas/simple-crud-chi/internal/model"
	"github.com/CesarSotnas/simple-crud-chi/internal/ports"
)

type championsRotationsRepository struct{}

func NewGetChampionsRotationsRepository() ports.ChampionsRotationRepositoryInterface {
	return &championsRotationsRepository{}
}

func (r *championsRotationsRepository) GetChampionsRotations() (model.ChampionsRotationResponse, error) {



	return data, nil
}
