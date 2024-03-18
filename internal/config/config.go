package config

import (
	"filmlib/internal/apperrors"
	"fmt"
	"os"
	"time"

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
	User              string `yaml:"-"`
	Password          string `yaml:"-"`
	Host              string `yaml:"-"`
	Port              uint64 `yaml:"port"`
	DBName            string `yaml:"-"`
	AppName           string `yaml:"-"`
	Schema            string `yaml:"schema"`
	ConnectionTimeout uint64 `yaml:"connection_timeout"`
}

type JWTConfig struct {
	Secret          string        `yaml:"-"`
	LifetimeSeconds uint          `yaml:"lifetime_seconds"`
	Lifetime        time.Duration `yaml:"-"`
}

type Config struct {
	API      *APIConfig      `yaml:"api"`
	Logging  *LoggingConfig  `yaml:"logging"`
	Server   *ServerConfig   `yaml:"server"`
	Database *DatabaseConfig `yaml:"db"`
	JWT      *JWTConfig      `yaml:"jwt"`
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

	config.Database.User, err = getDBUser()
	if err != nil {
		return nil, err
	}
	config.Database.Password, err = getDBPassword()
	if err != nil {
		return nil, err
	}
	config.Database.DBName, err = getDBName()
	if err != nil {
		return nil, err
	}
	config.Database.Host = getDBConnectionHost()

	config.JWT.Secret, err = getJWTSecret()
	if err != nil {
		return nil, err
	}
	config.JWT.Lifetime = time.Duration(config.JWT.LifetimeSeconds) * time.Second

	config.API.BaseUrl = config.API.BaseUrl + fmt.Sprint(config.API.APIVersion) + "/"

	return &config, nil
}

// getDBConnectionHost
// возвращает имя хоста из env для соединения с БД (по умолчанию localhost)
func getDBConnectionHost() string {
	host, hOk := os.LookupEnv("POSTGRES_HOST")
	if !hOk {
		return "localhost"
	}
	return host
}

// getDBConnectionHost
// возвращает пароль из env для соединения с БД
func getDBUser() (string, error) {
	user, uOk := os.LookupEnv("POSTGRES_USER")
	if !uOk {
		return "", apperrors.ErrDatabaseUserMissing
	}
	return user, nil
}

// getDBConnectionHost
// возвращает пароль из env для соединения с БД
func getDBPassword() (string, error) {
	pwd, pOk := os.LookupEnv("POSTGRES_PASSWORD")
	if !pOk {
		return "", apperrors.ErrDatabasePWMissing
	}
	return pwd, nil
}

// getDBConnectionHost
// возвращает пароль из env для соединения с БД
func getDBName() (string, error) {
	name, nOk := os.LookupEnv("POSTGRES_DB")
	if !nOk {
		return "", apperrors.ErrDatabaseNameMissing
	}
	return name, nil
}

func getJWTSecret() (string, error) {
	name, nOk := os.LookupEnv("JWT_SECRET")
	if !nOk {
		return "", apperrors.ErrJWTSecretMissing
	}
	return name, nil
}
