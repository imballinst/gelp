package gelp

import (
	"errors"
	"fmt"
	"strings"

	helpers "github.com/imballinst/gelp/src/helpers"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var Squash bool

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
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		fmt.Printf("You choose %q\n", result)
	},
}

func init() {
	// migrateCmd.Flags().BoolVarP(&Squash, "version", "v", false, "Show current gelp version")
}
