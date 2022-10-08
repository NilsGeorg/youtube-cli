package config

import (
	"crypto/md5" //nolint:gosec // not used for crypto
	"encoding/base64"
	"encoding/hex"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

type Authentication interface {
	GetToken() (string, error)
	SetToken(token string) error
}

type AuthenticationImpl struct{}

func (auth AuthenticationImpl) GetToken() (string, error) {
	path, err := auth.getTokenPath()
	if err != nil {
		return "", err
	}

	fileContent, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	decodeString, err := base64.StdEncoding.DecodeString(string(fileContent))
	if err != nil {
		return "", err
	}

	return string(decodeString), nil
}

func (auth AuthenticationImpl) SetToken(token string) error {
	path, err := auth.getTokenPath()
	if err != nil {
		return err
	}

	encodedToken := base64.StdEncoding.EncodeToString([]byte(token))

	var filePermission = 0o600

	return os.WriteFile(path, []byte(encodedToken), os.FileMode(filePermission))
}

func (auth AuthenticationImpl) getTokenPath() (string, error) {
	homeDir, err := homedir.Dir()

	if err != nil {
		return "", err
	}

	configDir := filepath.Join(homeDir, ".config", "youtube-cli")

	err = os.MkdirAll(configDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	tokenPath := filepath.Join(configDir, getConfigFileName())

	return tokenPath, nil
}

func getConfigFileName() string {
	hash := md5.Sum([]byte("youtube-cli-token")) //nolint:gosec // not used for crypto

	return hex.EncodeToString(hash[:]) + ".json"
}
