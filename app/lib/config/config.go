package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"regexp"
)

type AuthConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ReviewerGroup struct {
	BranchNameRegex *regexp.Regexp `yaml:"branch_name_regex,omitempty"`
	Reviewers       []string       `yaml:"reviewers"`
}

type Config struct {
	Authorization  AuthConfig      `yaml:"authorization"`
	ReviewerGroups []ReviewerGroup `yaml:"reviewer_groups"`
}

func GetConfig() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	file := path.Join(homeDir, ".config", "bbcli.yaml")
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func SaveConfig(config Config) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	file := path.Join(homeDir, ".config", "bbcli.yaml")
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
