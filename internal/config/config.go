package config

import (
	"filmlib/internal/apperrors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
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

type DatabaseConfig struct {
	User              string `yaml:"user"`
	Password          string `yaml:"-"`
	Host              string `yaml:"-"`
	Port              uint64 `yaml:"port"`
	DBName            string `yaml:"db_name"`
	AppName           string `yaml:"app_name"`
	Schema            string `yaml:"schema"`
	ConnectionTimeout uint64 `yaml:"connection_timeout"`
}

type Config struct {
	API      *APIConfig      `yaml:"api"`
	Logging  *LoggingConfig  `yaml:"logging"`
	Server   *ServerConfig   `yaml:"server"`
	Database *DatabaseConfig `yaml:"config"`
}

func LoadConfig(envPath string, configPath string) (*Config, error) {
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

	if envPath == "" {
		err = godotenv.Load()
	} else {
		err = godotenv.Load(envPath)
	}

	if err != nil {
		return nil, apperrors.ErrEnvNotFound
	}

	config.Database.Password, err = GetDBPassword()
	if err != nil {
		return nil, err
	}
	config.Database.Host = GetDBConnectionHost()

	config.API.BaseUrl = config.API.BaseUrl + fmt.Sprint(config.API.APIVersion) + "/"

	return &config, nil
}

// GetDBConnectionHost
// возвращает имя хоста из env для соединения с БД (по умолчанию localhost)
func GetDBConnectionHost() string {
	host, hOk := os.LookupEnv("POSTGRES_HOST")
	if !hOk {
		return "localhost"
	}
	return host
}

// getDBConnectionHost
// возвращает пароль из env для соединения с БД
func GetDBPassword() (string, error) {
	pwd, pOk := os.LookupEnv("POSTGRES_PASSWORD")
	if !pOk {
		return "", apperrors.ErrDatabasePWMissing
	}
	return pwd, nil
}
