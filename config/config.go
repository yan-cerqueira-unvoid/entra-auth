package config

import (
	"os"
)

type Config struct {
	Port           string
	MongoURI       string
	MongoDB        string
	JWTSecret      string
	JWTExpiryHours int
	EntraID        EntraIDConfig
}

type EntraIDConfig struct {
	TenantID     string
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

func NewConfig() *Config {
	jwtExpiryHours := 24 // Default to 24 hours

	return &Config{
		Port:           getEnv("PORT", "8080"),
		MongoURI:       getEnv("MONGO_URI", "mongodb://localhost:27017"),
		MongoDB:        getEnv("MONGO_DB", "auth_system"),
		JWTSecret:      getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpiryHours: jwtExpiryHours,
		EntraID: EntraIDConfig{
			TenantID:     getEnv("ENTRA_TENANT_ID", ""),
			ClientID:     getEnv("ENTRA_CLIENT_ID", ""),
			ClientSecret: getEnv("ENTRA_CLIENT_SECRET", ""),
			RedirectURI:  getEnv("ENTRA_REDIRECT_URI", "http://localhost:8080/api/auth/entra-callback"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
