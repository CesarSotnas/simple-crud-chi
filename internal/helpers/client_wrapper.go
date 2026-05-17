package helpers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/CesarSotnas/simple-crud-chi/internal/helpers"
)

func RequestBuilder[T any](endpoint string) (models.BaseResponse[T], error) {
	apiKey := os.Getenv("API_KEY")

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	url, err := mountURL(endpoint, apiKey)
	if err != nil {
		errMsg := errors.New(helpers.GenericError)
		return "", errMsg
	}

	response, err := makeTheRequest[T](client, url)

}

func mountURL(endpoint, apiKey string) (string, error) {
	endpointResponse, exists := helpers.ClientEndpoint[endpoint]
	if !exists {
		errMsg := errors.New(helpers.NotFound)
		return "", errMsg
	}

	url := endpointResponse + apiKey
	return url, nil
}

func makeTheRequest[T any](client http.Client, url string) (models.BaseResponse[T any], error) {
	response, err := client.Get(url)
	defer response.Body.Close()

	if err != nil {
		errMsg := errors.New(helpers.GenericError)
		return model.ChampionsRotationResponse{}, errMsg
	}

}