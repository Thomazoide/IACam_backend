package config

import "os"

type Config struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort string
	DBName string
}

func Load() *Config {
	return &Config{
		DBUser: getEnv("DBUSER"),
		DBPass: getEnv("DBPASS"),
		DBHost: getEnv("DBHOST"),
		DBPort: getEnv("DBPORT"),
		DBName: getEnv("DBNAME"),
	}
}

func getEnv(key string) string {
	return os.Getenv(key)
}
