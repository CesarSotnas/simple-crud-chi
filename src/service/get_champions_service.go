package service

import (
	"github.com/CesarSotnas/simple-crud-chi/src/internal"
	"github.com/CesarSotnas/simple-crud-chi/src/models"
)

type championsRotationsService struct {
	championsRotationsRepository internal.ChampionsRotationRepositoryInterface
}

func NewChampionsRotationsSerivce(championsRotationRepository internal.ChampionsRotationRepositoryInterface) internal.ChampionsRotationServiceInterface {
	return &championsRotationsService{
		championsRotationsRepository: championsRotationRepository,
	}
}

func (s *championsRotationsService) GetChampionsRotations() (models.ChampionsRotationResponse, error) {
	response, err := s.championsRotationsRepository.GetChampionsRotations()

	return response, err
}
