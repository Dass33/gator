package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func get_config_file_path() (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Can't find home directory: %v", err)
	}
	path = filepath.Join(path, configFileName)
	return path, nil
}

func write(config *Config, path string) error {
	data, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("Can't marshal to json: %v", err)
	}
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Can't open a config file: %v", err)
	}

	written, err := file.Write(data)
	if err != nil {
		return fmt.Errorf("Written %v bytes, writing to file had failed: %v", written, err)
	}

	return nil
}

func Read() (Config, error) {
	path, err := get_config_file_path()
	if err != nil {
		return Config{}, err
	}
	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("Can't open a config file: %v", err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return Config{}, fmt.Errorf("There was a problem with reading the config: %v", err)
	}

	var config Config
	json.Unmarshal(data, &config)
	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	path, err := get_config_file_path()
	if err != nil {
		return err
	}

	err = write(c, path)
	if err != nil {
		return fmt.Errorf("Can't set the username to config: %v", err)
	}
	return nil
}
