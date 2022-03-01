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
	Use:   "prune",
	Short: "Prunes the dangling remote branches in the local",
	Long: fmt.Sprintf(`Prune is used when you have a lot of dangling remote branches. When you delete a local branch,
Git still stores the "remote branch" locally (perhaps for safety measures). This can be checked with the command:

%s

There might be a lot of dangling remote branches (marked in red) that has been deleted in the remote Git repository.
"gelp prune" is a two-in-one command, which executes "git remote prune <remote> --dry-run" and "git remote prune <remote>".
gelp will also gives you a confirmation before deleting, so it acts as a "safeguard".
`, color.CyanString("git branch -a")),
	Example: fmt.Sprintf(`1) Prunes dangling remote "origin" branches
   %s

2) Prunes dangling remote "upstream" branches
   %s`,
		color.CyanString("gelp prune"),
		color.CyanString("gelp prune --origin upstream")),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return errors.New("`gelp prune` doesn't accept an argument. Please remove it")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		output, err := helpers.DoAndLog("prune", fmt.Sprintf("git remote prune %s --dry-run", pruneRemote))
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

		_, err = helpers.DoAndLog("prune", fmt.Sprintf("git remote prune %s", pruneRemote))
		if err != nil {
			panic(err)
		}
		helpers.Log("prune", "branch pruning completed")
	},
}

func init() {
	pruneCmd.Flags().StringVarP(&pruneRemote, "remote", "r", "origin", "The remote used to pull changes. Defaults to origin.")
}
