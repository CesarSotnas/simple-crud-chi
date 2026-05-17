package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CesarSotnas/simple-crud-chi/internal/helpers"
	"github.com/CesarSotnas/simple-crud-chi/internal/ports"
)

type championsRotationsController struct {
	championsRotationService ports.ChampionsRotationServiceInterface
}

func NewChampionsRotationsController(championsRotationService ports.ChampionsRotationServiceInterface) championsRotationsController {
	return championsRotationsController{
		championsRotationService: championsRotationService,
	}
}

func (c *championsRotationsController) GetChampionsController(w http.ResponseWriter, r *http.Request) {
	response, errResponse := c.championsRotationService.GetChampionsRotations()
	if errResponse != nil {

	}

	responseByte, err := json.Marshal(&response)
	if err != nil {
		w.Write([]byte(helpers.GenericError))
	}

	w.Write(responseByte)
}
