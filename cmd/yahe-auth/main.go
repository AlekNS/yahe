package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/alekns/yahe/internal/auth"
	cmds "github.com/alekns/yahe/internal/auth/presentation/commands"
)

var mainCommand = &cobra.Command{
	Use:     auth.ServiceName,
	Long:    auth.ServiceDesc,
	Version: auth.ServiceVersion}

func main() {
	viper.SetEnvPrefix("yahea")
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	mainFlags := mainCommand.PersistentFlags()
	mainFlags.String("config-file", "", "Config file")

	mainFlags.String("logging-level", "", "Default console level")
	viper.BindPFlag("logging.console.level", mainFlags.Lookup("logging-level"))

	mainCommand.AddCommand(cmds.RootCommand())

	if mainCommand.Execute() != nil {
		os.Exit(1)
	}
}
