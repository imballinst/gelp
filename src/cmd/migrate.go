package gelp

import (
	"errors"
	"fmt"
	"strings"

	helpers "github.com/imballinst/gelp/src/helpers"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var Base string

// Root command.
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate from one branch to another",
	Long: `Migrate is useful if you are working on a wrong branch.
Using it, you can move to the correct branch, along with your new changes. The changes in the
wrong branch will be removed.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("`git migrate` command needs 1 argument: target_branch")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		gitLog, err := helpers.ExecCommand("git", "log", "--oneline")
		if err != nil {
			panic(err)
		}

		prompt := promptui.Select{
			Label: "Select base commit",
			Items: strings.Split(gitLog, "\n"),
		}

		_, result, err := prompt.Run()

		if err != nil {
			panic(err)
		}

		currentBranch, err := helpers.ExecCommand("git", "rev-parse", "--abbrev-ref", "HEAD")
		if err != nil {
			panic(err)
		}

		fmt.Printf("Current branch: %q\n", currentBranch)
		fmt.Printf("You choose %q\n", result)
		out, err := helpers.ExecCommand("echo", currentBranch)
		if err != nil {
			panic(err)
		}

		fmt.Println(out)
		// _, err = helpers.ExecCommand("git")
	},
}

func init() {
	migrateCmd.Flags().StringVarP(&Base, "base", "b", "main", "The base branch used for the new branch")
}
