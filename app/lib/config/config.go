package config

import "regexp"

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
