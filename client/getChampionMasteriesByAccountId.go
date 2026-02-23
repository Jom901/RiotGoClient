package client

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type RewardConfigDto struct {
	RewardValue   string `json:rewardValue`
	RewardType    string `json:rewardType`
	MaximumReward int    `json:maximumReward`
}

type NextSeasonMilestoneDto struct {
	RequiredGradeCounts any             `json:requiredGradeCounts`
	RewardMarks         int             `json:rewardMarks`
	Bonus               bool            `json:bonus`
	RewardConfig        RewardConfigDto `json:rewardConfig`
}

type ChampionMasteryDto struct {
	AccountId                    string                 `json:puuid`
	ChampionPointsUntilNextLevel int64                  `json:championPointsUntilNextLevel`
	ChestGranted                 bool                   `json:chestGranted`
	ChampionId                   int64                  `json:championId`
	LastPlayTime                 int64                  `json:lastPlayTime`
	ChampionLevel                int                    `json:championLevel`
	ChampionPoints               int                    `json:championPoints`
	ChampionPointsSinceLastLevel int64                  `json:championPointsSinceLastLevel`
	MarkRequiredForNextLevel     int                    `json:markRequiredForNextLevel`
	ChampionSeasonMilestone      int                    `json:championSeasonMilestone`
	TokensEarned                 int                    `json:tokensEarned`
	MilestoneGrades              []string               `json:milestoneGrades`
	NextSeasonMilestone          NextSeasonMilestoneDto `json:nextSeasonMilestone`
}

func GetChampionMasteriesByAccountId(accountId string) ([]ChampionMasteryDto, error) {
	url := fmt.Sprintf("https://na1.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-puuid/%s?api_key=%s", accountId, apiKey)
	log.Println(url)
	resp, err := http.Get(url)
	var championMasteries []ChampionMasteryDto
	if err != nil {
		return championMasteries, fmt.Errorf("failed to get accounts: %w", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return championMasteries, fmt.Errorf("failed to read body: %w", err)
	}
	err = json.Unmarshal(body, &championMasteries)
	if err != nil {
		return championMasteries, fmt.Errorf("failed to unmarshal body: %w", err)
	}
	return championMasteries, nil
}
