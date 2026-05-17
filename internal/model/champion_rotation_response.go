package model

type ChampionsRotationResponse struct {
	FreeChampionId               []int `json:"freeChampionIds"`
	FreeChampionIdsForNewPlayers []int `json:"freeChampionIdsForNewPlayers"`
	MaxNewPlayerLevel            int   `json:"maxNewPlayerLevel"`
}

type BaseResponse[T any] interface {
}
