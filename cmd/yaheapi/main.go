package main

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/alekns/yahe/internal"
	cmds "github.com/alekns/yahe/internal/presentation/commands"
)

var mainCommand = &cobra.Command{
	Use:     internal.ServiceName,
	Short:   "Hyperledger Explorer",
	Long:    "Yet Another Hyperledger (fabric) Explorer",
	Version: internal.ServiceVersion}

func main() {
	viper.SetEnvPrefix("yahe")
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
