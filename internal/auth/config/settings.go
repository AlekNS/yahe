package config

import (
	"github.com/spf13/viper"

	"github.com/alekns/yahe/internal/config"
)

// Settings gathering all application settings.
type Settings struct {
	Logger *config.LoggerSettings
	Jwt    *JwtSettings
	Users  *UsersSettings
}

// GetSettings reads all from config file and env.
func GetSettings(viper *viper.Viper) *Settings {
	return &Settings{
		Logger: config.FromViperLoggerSettings(viper),
		Jwt:    jwtSettingsGetAndValidate(viper),
		Users:  usersSettingsGetAndValidate(viper),
	}
}
