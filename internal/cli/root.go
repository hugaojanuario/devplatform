package cli

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "devplatform",
	Short: "...",
	Long:  "...",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(healthCmd)
}
