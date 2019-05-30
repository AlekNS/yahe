package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "auth",
	Short: "Auth commands",
}

// RootCommand .
func RootCommand() *cobra.Command {
	rootCmd.AddCommand(serveCmd)
	return rootCmd
}
