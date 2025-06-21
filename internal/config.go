package internal

import (
	"encoding/json"
	"os"
)

type Config struct {
	Tickers        []string `json:"tickers"`
	Interval       string   `json:"interval"`
	Range          string   `json:"range"`
	RefreshMinutes int      `json:"refresh_minutes"`
	APIKey         string   `json:"api_key"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
