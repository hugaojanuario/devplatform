package cli

import (
	"fmt"

	"github.com/hugaojanuario/devplatform/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show DevPlatform version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DevPlatform version", version.Version)
	},
}
