package config

import (
	"encoding/json"
	"os"
	"path"
	"regexp"
)

type AuthConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReviewerGroup struct {
	BranchNameRegex *regexp.Regexp `json:"branch_name_regex,omitempty"`
	Reviewers       []string       `json:"reviewers"`
}

type Config struct {
	Authorization  AuthConfig      `json:"authorization"`
	ReviewerGroups []ReviewerGroup `json:"reviewer_groups"`
}

func GetConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	file := path.Join(homeDir, ".config", "bbcli.json")
	data, err := os.ReadFile(file)
	if err != nil {
		return &Config{}, nil
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func SaveConfig(config *Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	file := path.Join(homeDir, ".config", "bbcli.json")
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, data, 0600)
	if err != nil {
		return err
	}
	return nil
}
