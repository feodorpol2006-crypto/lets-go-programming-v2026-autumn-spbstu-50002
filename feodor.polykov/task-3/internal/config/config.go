package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	InputPath  string `yaml:"input-file"`
	OutputPath string `yaml:"output-file"`
}

func ReadConfig(filePath string) (*AppConfig, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read config file: %w", err)
	}

	var appCfg AppConfig

	if err := yaml.Unmarshal(fileData, &appCfg); err != nil {
		return nil, fmt.Errorf("could not parse YAML data: %w", err)
	}

	if appCfg.InputPath == "" || appCfg.OutputPath == "" {
		return nil, errors.New("both input-file and output-file must be provided in config")
	}

	return &appCfg, nil
}
