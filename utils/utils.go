package utils

import (
	"encoding/json"
	"github.com/CarosDrean/api-amachay/models"
	"os"
)

func GetConfiguration() (models.Configuration, error) {
	config := models.Configuration{}
	file, err := os.Open("./configuration.json")

	if err != nil {
		return config, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		return config, err
	}

	return config, nil
}

func GetConfigurationTelegram() (models.ConfigurationTelegram, error) {
	config := models.ConfigurationTelegram{}
	file, err := os.Open("./configuration-telegram.json")

	if err != nil {
		return config, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		return config, err
	}

	return config, nil
}
