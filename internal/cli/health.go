package cli

import (
	"github.com/hugaojanuario/devplatform/internal/client"
	"github.com/hugaojanuario/devplatform/pkg/config"
	"github.com/spf13/cobra"
)

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Show DevPlatform health",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := config.LoadEnvFile()

		err := client.Client(cfg.Api.Url + "/health")
		if err != nil {
			cmd.PrintErrln("Api offline")
		} else {
			cmd.Println("Api online")
		}
	},
}
