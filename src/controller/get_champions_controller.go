package controller

import (
	"encoding/json"
	"net/http"

	"github.com/CesarSotnas/simple-crud-chi/src/helpers"
	"github.com/CesarSotnas/simple-crud-chi/src/internal"
)

type championsRotationsController struct {
	championsRotationService internal.ChampionsRotationServiceInterface
}

func NewChampionsRotationsController(championsRotationService internal.ChampionsRotationServiceInterface) championsRotationsController {
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
