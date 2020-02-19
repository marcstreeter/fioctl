package cmd

import (
	"github.com/spf13/cobra"
)

var keysCmd = &cobra.Command{
	Use:              "keys",
	Short:            "Manage keys in use by your factory fleet",
	PersistentPreRun: assertLogin,
}

func init() {
	rootCmd.AddCommand(keysCmd)
}