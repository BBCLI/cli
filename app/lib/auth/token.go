package auth

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func GetToken() (string, error) {
	homeDir, err := os.UserHomeDir()
	tokenFilePath := filepath.Join(homeDir, ".config", "bbcli.token")
	file, err := os.ReadFile(tokenFilePath)
	if err != nil {
		return "", err
	}
	fmt.Println(file)
	token := string(file)
	if token == "" {
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

	_, err = file.Write([]byte(token))
	if err != nil {
		return err
	}
	return nil
}
