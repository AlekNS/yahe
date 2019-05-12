package helpers

import (
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetSettings .
func GetSettings(serviceName string, cmd *cobra.Command, v *viper.Viper) *viper.Viper {
	var flagConfig = cmd.Flag("config-file")

	if len(flagConfig.Value.String()) == 0 {
		// if config not specified
		v.SetConfigName("config")
		v.AddConfigPath(filepath.FromSlash("/etc/" + serviceName))
		v.AddConfigPath(filepath.FromSlash("$HOME/." + serviceName))
		v.AddConfigPath(".")
	} else {
		v.SetConfigFile(flagConfig.Value.String())
	}

	var err = v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return v
}
