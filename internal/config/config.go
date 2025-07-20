package config

import (
	"encoding/json"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	UrlDb string `json:"db_url"`
	UserName string `json:"username"`
}

func Read() (Config, error) {
	
	config_path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	json_str, err := os.ReadFile(config_path)
	if err != nil {
		return Config{}, err
	}
	var the_config Config
	if err = json.Unmarshal([]byte(json_str), &the_config); err != nil {
		return the_config, err
	}
	return the_config, nil
}

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	config_path := home_dir + configFileName
	return config_path, nil
}

func (c *Config) SetUser(username string) error {
	c.UserName = username

	config_path, err := getConfigFilePath()
	if err != nil {
		return err
	} 

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(config_path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}