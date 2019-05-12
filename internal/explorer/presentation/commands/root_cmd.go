package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "Explorer Api Commands",
}

// RootCommand .
func RootCommand() *cobra.Command {
	return rootCmd
}
