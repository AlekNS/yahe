package config

import (
	"github.com/spf13/viper"

	"github.com/alekns/yahe/internal/config"
)

// Settings gathering all application settings.
type Settings struct {
	Logger *config.LoggerSettings
	HTTP   *config.HTTPSettings
	Jwt    *JwtSettings
	Users  *UsersSettings
}

// GetSettings reads all from config file and env.
func GetSettings(viper *viper.Viper) *Settings {
	return &Settings{
		Logger: config.FromViperLoggerSettings(viper),
		HTTP:   config.FromViperHTTPSettings(viper),
		Jwt:    jwtSettingsGetAndValidate(viper),
		Users:  usersSettingsGetAndValidate(viper),
	}
}
