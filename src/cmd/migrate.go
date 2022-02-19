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
What this command does is:

- Create a new branch if the target branch doesn't exist, or 
  create a copy of that branch with the name "backup--{branch_name}"
- Checkout to that branch
- Cherry pick the selected commit(s)
- Go back to the old branch

The old branch will not be touched. The wrong commits you can resolve yourselves
using "git rebase" or "git reset", depending on the scenario.`,
	Example: fmt.Sprintf(`1) Migrate to "test-branch" using base branch "main"
   %s

2) Migrate to "hotfix" using base branch "dev"
   %s`, helpers.GetBlueText("gelp migrate test-branch"), helpers.GetBlueText("gelp migrate hotfix --base dev")),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("`git migrate` command needs 1 argument: target_branch")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		gitLog, err := helpers.ExecCommand("git log --oneline")
		if err != nil {
			panic(err)
		}

		// Select the start commit.
		prompt := promptui.Select{
			Label: "Select start commit to migrate",
			Items: strings.Split(gitLog, "\n"),
		}

		_, startCommit, err := prompt.Run()
		if err != nil {
			panic(err)
		}

		// Select the end commit.
		prompt = promptui.Select{
			Label: "Select end commit to migrate",
			Items: strings.Split(gitLog, "\n"),
		}

		_, endCommit, err := prompt.Run()
		if err != nil {
			panic(err)
		}

		helpers.Migrate(args[0], Base, startCommit, endCommit)
	},
}

func init() {
	migrateCmd.Flags().StringVarP(&Base, "base", "b", "main", "The base branch used for the new branch")
}
