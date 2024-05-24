package auth

import (
	"errors"
	"os"
	"path/filepath"
)

func GetToken() (string, error) {
	var file, err = os.ReadFile("~/.config/bbcli.token")
	if err != nil {
		return "", err
	}
	var token = string(file)
	if token != "" {
		return "", errors.New("no token")
	}
	return token, nil
}

func SetToken(token string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	tokenFilePath := filepath.Join(homeDir, ".config", "bbcli.token")
	file, err := os.Create(tokenFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, writeErr := file.Write([]byte(token))
	if writeErr != nil {
		return writeErr
	}
	return nil
}
