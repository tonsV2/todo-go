package configuration

import (
	"log"
	"os"
	"strconv"
)

// Inspiration: https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f

func ProvideConfiguration() Configuration {
	return Configuration{
		BaseUrl: requireEnv("BASE_URL"),
		Authentication: authentication{
			Keys: keys{
				PrivateKey: requireEnv("PRIVATE_KEY_FILE"),
				PublicKey:  requireEnv("PUBLIC_KEY_FILE"),
			},
			RefreshTokenSecretKey:         requireEnv("REFRESH_TOKEN_SECRET_KEY"),
			AccessTokenExpirationSeconds:  getEnvAsInt("ACCESS_TOKEN_EXPIRATION_IN_SECONDS"),
			RefreshTokenExpirationSeconds: getEnvAsInt("REFRESH_TOKEN_EXPIRATION_IN_SECONDS"),
		},
		Postgresql: postgresql{
			Host:         requireEnv("DATABASE_HOST"),
			Port:         getEnvAsInt("DATABASE_PORT"),
			Username:     requireEnv("DATABASE_USERNAME"),
			Password:     requireEnv("DATABASE_PASSWORD"),
			DatabaseName: requireEnv("DATABASE_NAME"),
		},
		Redis: redis{
			Host: requireEnv("REDIS_HOST"),
			Port: getEnvAsInt("REDIS_PORT"),
		},
		AdminUser: user{
			Email:    requireEnv("ADMIN_USER_EMAIL"),
			Password: requireEnv("ADMIN_USER_PASSWORD"),
		},
		DefaultUser: user{
			Email:    requireEnv("USER_EMAIL"),
			Password: requireEnv("PASSWORD"),
		},
	}
}

type Configuration struct {
	BaseUrl        string
	Authentication authentication
	Postgresql     postgresql
	Redis          redis
	AdminUser      user
	DefaultUser    user
}

type authentication struct {
	Keys                          keys
	RefreshTokenSecretKey         string
	AccessTokenExpirationSeconds  int
	RefreshTokenExpirationSeconds int
}

type keys struct {
	PrivateKey string
	PublicKey  string
}

type postgresql struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}

type redis struct {
	Host string
	Port int
}

type user struct {
	Email    string
	Password string
}

func requireEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatalf("Can't find environment varialbe: %s\n", key)
	}
	return value
}

func getEnvAsInt(key string) int {
	valueStr := requireEnv(key)
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Fatalf("Can't parse value as integer: %s", err.Error())
	}
	return value
}
