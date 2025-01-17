package config

import (
	"encoding/json"
	"log"
)

// TrelloConfig Config for trello connector
type TrelloConfig struct {
	rawConfig rawTrelloConfig
}

type rawTrelloConfig struct {
	BaseURL  string `json:"baseUrl"`
	BoardAPI string `json:"boardApi"`
}

// NewTrelloConfig Create trello config instance
func NewTrelloConfig() *TrelloConfig {
	const configName = "trello.json"
	const envName = "ENV"

	configHelper := NewHelper(configName, envName)
	trelloConfig := &TrelloConfig{
		rawConfig: rawTrelloConfig{},
	}
	error := json.Unmarshal(configHelper.GetConfig(), &trelloConfig.rawConfig)

	if error != nil {
		log.Fatalf("Error unmarshalling config with name %s %v", configName, error)
	}

	return trelloConfig
}

// GetBaseURL Return base url from config
func (trelloConfig *TrelloConfig) GetBaseURL() string {
	return trelloConfig.rawConfig.BaseURL
}

// GetBoardAPIURL Return trello board api url
func (trelloConfig *TrelloConfig) GetBoardAPIURL() string {
	return trelloConfig.rawConfig.BoardAPI
}
