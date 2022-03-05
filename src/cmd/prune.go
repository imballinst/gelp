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

		fmt.Printf("\n%s\n\n", color.CyanString(output))

		validate := func(s string) error {
			if len(s) == 1 && strings.Contains("YyNn", s) || len(s) == 0 {
				return nil
			}
			return errors.New("invalid input")
		}
		prompt := promptui.Prompt{
			Label:     "Prune the listed branches above",
			IsConfirm: true,
			Validate:  validate,
			Default:   "y",
		}
		result, err := prompt.Run()
		isErrAbort := errors.Is(err, promptui.ErrAbort)

		if err != nil && !isErrAbort {
			// Abort error only.
			panic(err)
		}

		if strings.ToLower(result) == "n" {
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
