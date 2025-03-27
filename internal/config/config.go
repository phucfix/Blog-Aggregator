package config

import (
    "os"
    "path/filepath"
    "encoding/json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
    DBURL           string `json:"db_url"`
    CurrentUsername string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
    c.CurrentUsername = userName
    return write(*c);
}

// Read a json file
func Read() (Config, error) {
    fullPath, err := getConfigFilePath()
    if err != nil {
        return Config{}, err
    }

    file, err := os.Open(fullPath)
    if err != nil {
        return Config{}, err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    cfg := Config{}
    if err := decoder.Decode(&cfg); err != nil {
        return Config{}, err
    }

    return cfg, nil
}

// Write a json struct to file
func write(cfg Config) error {
    fullPath, err := getConfigFilePath()
    if err != nil {
        return err
    }

    file, err := os.Create(fullPath)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := json.NewEncoder(file)
    if err = encoder.Encode(cfg); err != nil {
        return err
    }

    return nil
}

func getConfigFilePath() (string, error) {
    home , err := os.UserHomeDir()
    if err != nil {
        return "", err
    }

    return filepath.Join(home, configFileName), nil
}

