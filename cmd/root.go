package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gobox",
	Short: "Manage and reuse Go packages easily",
	Long: `Gobox helps you fetch, track, and reuse Go packages easily.
It provides commands like 'get', 'list', 'remove', and 'init' to streamline your Go project workflow.`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		return err
	}
	return nil
}

// func init() {
// 	storage.Init()
// }
