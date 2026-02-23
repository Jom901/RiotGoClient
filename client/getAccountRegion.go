package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AccountRegionDTO struct {
	AccountId string `json:puuid`
	Game      string `json:game`
	Region    string `json:region`
}

func GetAccountRegionByGame(accountId string, game string) (AccountRegionDTO, error) {
	url := fmt.Sprintf("https://americas.api.riotgames.com/riot/account/v1/region/by-game/%s/by-puuid/%s?api_key=%s", game, accountId, apiKey)
	log.Println(url)
	resp, err := http.Get(url)

	var account AccountRegionDTO
	if err != nil {
		return account, fmt.Errorf("failed to get account: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return account, fmt.Errorf("failed to read body: %w", err)
	}
	err = json.Unmarshal(body, &account)
	if err != nil {
		return account, fmt.Errorf("failed to unmarshal body: %w", err)
	}
	return account, nil
}
