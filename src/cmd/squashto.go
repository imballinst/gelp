package cmd

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	helpers "github.com/imballinst/gelp/src/helpers"
	"github.com/spf13/cobra"
)

var squashNewBaseBranch string

// Root command.
var squashToCmd = &cobra.Command{
	Use:   "squashto",
	Short: "Squashes your changes and moves it to a new branch",
	Long: `Squashto is used usually when you have finished on a development branch and have opened a PR,
but the PR has not been reviewed and you need to use the current changes to work on another task. We can
just checkout to a new branch... although that would bring all of the commits with us (which can contain noise).

This command is a shorthand for:

- Checkout to an existing branch (or create a new one if it doesn't exist)
- Squash merge from the old branch

As an important note, "gelp squash" doesn't automatically resolve conflicts.`,
	Example: fmt.Sprintf(`1) Squashes the current changes to "test-branch" using base branch "main"
   %s

2) Squashes the current changes to "hotfix" using base branch "dev"
   %s`,
		color.CyanString("gelp squashto test-branch"),
		color.CyanString("gelp squashto hotfix --base dev")),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("`gelp squashto` command needs 1 argument: base_branch")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Squashto(args[0], squashNewBaseBranch)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	squashToCmd.Flags().StringVarP(&squashNewBaseBranch, "base", "b", "main", "The base branch used for the new branch. Defaults to main.")
}
