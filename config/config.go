package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/cetorres/ai-buddy/util"
)

type Config struct {
	GoogleAPIKey string `json:"googleApiKey"`
	OpenAIAPIKey string `json:"openaiApiKey"`
}

func makeConfig(googleApiKey string, openaiApiKey string) Config {
	return Config{
		GoogleAPIKey: googleApiKey,
		OpenAIAPIKey: openaiApiKey,
	}
}

func GetConfigDirectory() string {
	const configDir = "/.config/ai-buddy/"
	userConfigDir, err := os.UserHomeDir()
	if err != nil {
		util.PrintError(fmt.Sprintf("Could not get user home directory: %s,", err))
		return "." + configDir
	}
	return userConfigDir + configDir
}

func getConfigFilePath() string {
	return GetConfigDirectory() + "config.json"
}

func GetGoogleAPIKey() string {
	config := GetConfig()
	return config.GoogleAPIKey
}

func GetOpenAIAPIKey() string {
	config := GetConfig()
	return config.OpenAIAPIKey
}

func createConfigFile() (*os.File, error) {
	configFilePath := getConfigFilePath()

	if err := os.MkdirAll(filepath.Dir(configFilePath), 0755); err != nil {
			return nil, err
	}
	return os.Create(configFilePath)
}

func GetConfig() Config {
	config := makeConfig("", "")
	var file *os.File

	file, err := os.OpenFile(getConfigFilePath(), os.O_RDWR, 0644)
	if err != nil {
		file, err = createConfigFile()
		if err != nil {
			util.PrintError(fmt.Sprintf("Could not create config file: %s,", err))
			return config
		}
		SetConfig(config)
		return config
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		util.PrintError(fmt.Sprintf("Could not read contents of config file: %s,", err))
		return config
	}

	if err := json.Unmarshal(fileBytes, &config); err != nil {
		SetConfig(config)
		return config
	}

	return config
}

func SetConfig(config Config) bool {
	bytes, _ := json.MarshalIndent(config, "", "\t")
	if err := os.WriteFile(getConfigFilePath(), bytes, 0644); err != nil {
		util.PrintError(fmt.Sprintf("Could not save config file: %s,", err))
		return false
	}
	return true
}