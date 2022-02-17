package gelp

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show current version",
	Long:  `gelp is maintained using semantic versioning. By learning about the version, you can crosscheck the current bugs and features.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Use something like git describe --tags $(git rev-list --tags --max-count=1).
		// Source: https://stackoverflow.com/questions/1404796/how-can-i-get-the-latest-tag-name-in-current-branch-in-git.
		// Perhaps we should have a JSON or something.
		fmt.Println("test22")
	},
}
