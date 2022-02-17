package gelp

import (
	"fmt"
	"os"

	helpers "github.com/imballinst/gelp/src/helpers"
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
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		content, err := helpers.ReadFileContent(wd + "/version.txt")
		if err != nil {
			panic(err)
		}

		fmt.Println(content)
	},
}
