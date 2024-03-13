package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type APIConfig struct {
	BaseUrl    string `yaml:"base_url"`
	APIVersion uint   `yaml:"version"`
}

type ServerConfig struct {
	Port uint `yaml:"port"`
}

type LoggingConfig struct {
	Level                  string `yaml:"level"`
	DisableTimestamp       bool   `yaml:"disable_timestamp"`
	FullTimestamp          bool   `yaml:"full_timestamp"`
	DisableLevelTruncation bool   `yaml:"disable_level_truncation"`
	LevelBasedReport       bool   `yaml:"level_based_report"`
	ReportCaller           bool   `yaml:"report_caller"`
}

type Config struct {
	API     *APIConfig     `yaml:"api"`
	Logging *LoggingConfig `yaml:"logging"`
	Server  *ServerConfig  `yaml:"server"`
}

func LoadConfig(configPath string) (*Config, error) {
	var (
		config Config
		err    error
	)

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	config.API.BaseUrl = config.API.BaseUrl + fmt.Sprint(config.API.APIVersion) + "/"

	return &config, nil
}
