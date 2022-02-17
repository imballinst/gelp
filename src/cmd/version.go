package gelp

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Long:  `gelp is maintained using semantic versioning. By learning about the version, you can crosscheck the current bugs and features.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(os.Getenv("version"))
	},
}
