package config

import (
	"github.com/spf13/viper"
)

// HTTPSettings .
type HTTPSettings struct {
	// Bind defines binding interface and port.
	Bind string

	// LogAccess declares of logging of entrypoint access.
	LogAccess bool
}

// FromViperHTTPSettings fill-up configuration structure from viper.
func FromViperHTTPSettings(v *viper.Viper) *HTTPSettings {
	const (
		khttpbind = "endpoints.http.bind"
		khttplog  = "endpoints.http.logaccess"
	)

	// Setup defaults
	v.SetDefault(khttpbind, ":8080")
	v.SetDefault(khttplog, false)

	// Validations
	// level := v.GetString(khttpbind)
	// switch level {
	// case "debug", "info", "warn", "error":
	// default:
	// 	panic(kconsolelevel + " has invalid value")
	// }

	return &HTTPSettings{
		Bind:      v.GetString(khttpbind),
		LogAccess: v.GetBool(khttplog),
	}
}
