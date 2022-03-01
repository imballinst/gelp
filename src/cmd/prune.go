package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	helpers "github.com/imballinst/gelp/src/helpers"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var pruneRemote string

// Root command.
var pruneCmd = &cobra.Command{
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

2) Update "dev" branch after merging a PR to "dev" in remote repository
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
		output, err := helpers.DoAndLog("prune", fmt.Sprintf("git prune remote %s --dry-run", pruneRemote))
		if err != nil {
			panic(err)
		}

		prompt := promptui.Prompt{
			Label: fmt.Sprintf("The following branches will be deleted:\n%s\n%s",
				color.CyanString(output),
				"Type \"y/Y\" to proceed deleting then press Enter, else \"N/n\""),
		}
		result, err := prompt.Run()
		if err != nil {
			panic(err)
		}

		if result != "" || strings.ToLower(result) != "y" {
			helpers.Log("prune", "branch pruning cancelled")
			return
		}

		_, err = helpers.DoAndLog("prune", fmt.Sprintf("git prune remote %s", pruneRemote))
		if err != nil {
			panic(err)
		}
		helpers.Log("prune", "branch pruning completed")
	},
}

func init() {
	pruneCmd.Flags().StringVarP(&pruneRemote, "remote", "r", "origin", "The remote used to pull changes. Defaults to origin.")
}
