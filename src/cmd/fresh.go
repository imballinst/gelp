package cmd

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	helpers "github.com/imballinst/gelp/src/helpers"
	"github.com/spf13/cobra"
)

var freshBaseBranch string

// Root command.
var freshCmd = &cobra.Command{
	Use:   "fresh",
	Short: "Checkout to a new branch using the base branch as HEAD",
	Long: `Fresh is used to, as its name suggest, start fresh. This usually happens
after we have finished working on a branch and we want to work on another unrelated task.

On such a case, we want to make sure that we are using the base branch as HEAD because
otherwise the commits from the previously worked branch will be taken along to this new branch.`,
	Example: fmt.Sprintf(`1) Start fresh to "hotfix-1" branch using "main" as HEAD
   %s

2) Start fresh to "hotfix-1" branch using "dev" as HEAD
   %s`,
		color.CyanString("gelp fresh hotfix"),
		color.CyanString("gelp fresh hotfix --base dev")),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("`gelp fresh` command needs 1 argument: base_branch")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Fresh(args[0], freshBaseBranch)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	freshCmd.Flags().StringVarP(&freshBaseBranch, "base", "b", "main", "The base branch used for the new branch. Defaults to main.")
}
