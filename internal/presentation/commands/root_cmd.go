package cmds

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "explorer",
	Short: "Explorer commands",
}

// RootCommand .
func RootCommand() *cobra.Command {
	return rootCmd
}
