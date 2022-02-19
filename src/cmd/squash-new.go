package gelp

import (
	"errors"
	"fmt"
	"strings"

	"github.com/fatih/color"
	helpers "github.com/imballinst/gelp/src/helpers"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var SquashNewBaseBranch string

// Root command.
var squashNewCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate from one branch to another",
	Long: `Migrate is useful if you are working on a wrong branch.
What this command does is:

- Create a new branch if the target branch doesn't exist, or 
  create a copy of that branch with the name "backup--{branch_name}"
- Checkout to that branch
- Cherry pick the selected commit(s)

The old branch will not be touched. The wrong commits you can resolve yourselves
using "git rebase" or "git reset", depending on the scenario. As an important note,
"gelp migrate" doesn't automatically resolve conflicts.`,
	Example: fmt.Sprintf(`1) Migrate to "test-branch" using base branch "main"
   %s

2) Migrate to "hotfix" using base branch "dev"
   %s`, color.BlueString("gelp migrate test-branch"), color.BlueString("gelp migrate hotfix --base dev")),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("`git migrate` command needs 1 argument: target_branch")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO(imballinst): this will be used later.
		// currentBranchOutput, err := ExecCommand("git rev-parse --abbrev-ref HEAD")
		gitLog, err := helpers.ExecCommand("git log --oneline")
		if err != nil {
			panic(err)
		}

		// Select the start commit.
		gitLogArray := strings.Split(gitLog, "\n")
		prompt := promptui.Select{
			Label: "Select start commit to migrate",
			Items: gitLogArray,
		}

		_, startCommit, err := prompt.Run()
		if err != nil {
			panic(err)
		}

		// Colorize the selected commit.
		endGitLogArray := gitLogArray

		for i, commit := range endGitLogArray {
			if commit == startCommit {
				endGitLogArray[i] = color.New(color.BgWhite).Sprint(commit)
			}
		}

		// Select the end commit.
		prompt = promptui.Select{
			Label: "Select end commit to migrate",
			Items: endGitLogArray,
		}

		_, endCommit, err := prompt.Run()
		if err != nil {
			panic(err)
		}

		helpers.Migrate(args[0], SquashNewBaseBranch, startCommit, endCommit)
	},
}

func init() {
	squashNewCmd.Flags().StringVarP(&SquashNewBaseBranch, "base", "b", "main", "The base branch used for the new branch")
}