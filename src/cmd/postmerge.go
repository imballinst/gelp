package cmd

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	helpers "github.com/imballinst/gelp/src/helpers"
	"github.com/spf13/cobra"
)

var postMergeBaseBranch string
var postMergeRemote string

// Root command.
var postMergeCmd = &cobra.Command{
	Use:   "postmerge",
	Short: "Go to the base branch, pulls the latest changes, then deletes the previous branch",
	Long: `Postmerge is used after you have merged a pull request in the remote. In that scenario,
the local branch doesn't matter anymore, hence what we usually do is:

- Checkout to the base branch with "git checkout {base}"
- Pull the latest changes with "git pull {remote} {base}"
- Delete the branch that was merged just now with "git branch -D {base}"

"gelp postmerge" is a shorthand to do all 3 of these.`,
	Example: fmt.Sprintf(`1) Update "main" branch after merging a PR to "main" in remote repository
   %s

2) Update "dev" branch after merging a PR to "dev" in remote forked repository
   %s

3) Update "dev" branch after merging a PR to "dev" in remote repository
   %s`,
		color.CyanString("gelp postmerge"),
		color.CyanString("gelp postmerge --base dev"),
		color.CyanString("gelp postmerge --base dev --remote upstream")),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("`gelp postmerge` doesn't accept an argument. Please remove it")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := helpers.Postmerge(postMergeRemote, postMergeBaseBranch)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	postMergeCmd.Flags().StringVarP(&postMergeBaseBranch, "base", "b", "main", "The base branch used for the new branch. Defaults to main.")
	postMergeCmd.Flags().StringVarP(&postMergeRemote, "remote", "r", "origin", "The remote used to pull changes. Defaults to origin.")
}
