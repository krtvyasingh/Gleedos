package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DefaultQuality string `json:"default_quality"`
	DefaultOutput  string `json:"default_output"`
	Threads        int    `json:"threads"`
	Concurrency    int    `json:"concurrency"`
	RateLimit      string `json:"rate_limit"`
	AutoResume     bool   `json:"auto_resume"`
}

func DefaultConfig() *Config {
	home, _ := os.UserHomeDir()
	return &Config{
		DefaultQuality: "1080p",
		DefaultOutput:  filepath.Join(home, "Downloads", "Gleedos"),
		Threads:        4,
		Concurrency:    3,
		AutoResume:     true,
	}
}

func Load(path string) (*Config, error) {
	cfg := DefaultConfig()
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return cfg, nil
		}
		path = filepath.Join(home, ".gleedos", "config.json")
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}
	if err := json.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (c *Config) Save(path string) error {
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		path = filepath.Join(home, ".gleedos", "config.json")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}
