package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type AccountDto struct {
	AccountId string `json:"puuid"`
	GameName  string `json:"gameName"`
	TagLine   string `json:"tagLine"`
}

func GetAccountByGameName(gameName string, tagLine string) (AccountDto, error) {
	url := fmt.Sprintf("https://americas.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s", gameName, tagLine, apiKey)
	log.Println(url)
	resp, err := http.Get(url)
	var account AccountDto
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
