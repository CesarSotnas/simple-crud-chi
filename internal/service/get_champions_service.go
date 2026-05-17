package service

import (
	"github.com/CesarSotnas/simple-crud-chi/internal/model"
	"github.com/CesarSotnas/simple-crud-chi/internal/ports"
)

type championsRotationsService struct {
	championsRotationsRepository ports.ChampionsRotationRepositoryInterface
}

func NewChampionsRotationsSerivce(championsRotationRepository ports.ChampionsRotationRepositoryInterface) ports.ChampionsRotationServiceInterface {
	return &championsRotationsService{
		championsRotationsRepository: championsRotationRepository,
	}
}

func (s *championsRotationsService) GetChampionsRotations() (model.ChampionsRotationResponse, error) {
	response, err := s.championsRotationsRepository.GetChampionsRotations()

	return response, err
}
