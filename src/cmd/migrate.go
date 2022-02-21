package gelp

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/fatih/color"
	helpers "github.com/imballinst/gelp/src/helpers"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// Flag types.
// Using camel-case because this is local.
type MigrateModeType int

const (
	MigrateModeSingle MigrateModeType = iota + 1
	MigrateModeMultiple
	MigrateModeRange
)

// Flags.
var migrateBaseBranch string
var migrateMode string
var migrateWithDate bool

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

The old branch will not be touched. The wrong commits you can resolve yourselves
using "git rebase" or "git reset", depending on the scenario. As an important note,
"gelp migrate" doesn't automatically resolve conflicts.`,
	Example: fmt.Sprintf(`1) Migrate to "test-branch" using base branch "main"
   %s

2) Migrate to "hotfix" using base branch "dev"
   %s`, color.BlueString("gelp migrate test-branch"), color.BlueString("gelp migrate hotfix --base dev")),
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("`gelp migrate` command needs 1 argument: target_branch")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Validate flag.
		parsedMode, err := parseMigrateMode(migrateMode)
		if err != nil {
			panic(err)
		}

		// Get list of commits.
		gitLog, err := helpers.ExecCommand("git log --oneline")
		if err != nil {
			panic(err)
		}
		gitLogArray := helpers.ExtractRevisionAndTitleFromCommits(strings.Split(gitLog, "\n"), migrateWithDate)

		// Use prompt to get inputs, then resolve the commits.
		revisions := ResolveRevisionsFromMigratePrompt(gitLogArray, parsedMode, nil)

		// Migrate.
		err = helpers.Migrate(args[0], migrateBaseBranch, revisions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	migrateCmd.Flags().StringVarP(&migrateBaseBranch, "base", "b", "main", "The base branch used for the new branch")
	migrateCmd.Flags().StringVarP(&migrateMode, "mode", "m", "single", "The commit picking mode: single, multiple, or range. Defaults to single.")
	migrateCmd.Flags().BoolVarP(&migrateWithDate, "date", "d", false, "Show the date of each commit when picking commit(s).")
}

// Semi-helper functions.
func ResolveRevisionsFromMigratePrompt(gitLogArray []string, parsedMode MigrateModeType, reader io.ReadCloser) []string {
	// Start picking commits.
	var pickedIndexes []int

	switch parsedMode {
	case MigrateModeMultiple:
		{
			// If it is "multiple", we append the "Finish picking commits" option.
			gitLogArray = append([]string{"-- Finish picking commits --"}, gitLogArray...)
			prompt := promptui.Select{
				Label: "Select commits to migrate (picked commits will be reordered by date)",
				Items: gitLogArray,
				Stdin: reader,
			}

			// Repeat picking commits until the first index is chosen.
			var pickedIndex int = -1

			for pickedIndex != 0 {
				pickedIndex, _, err := prompt.Run()
				if err != nil {
					panic(err)
				}

				if pickedIndex != 0 {
					// Only append if the picked index is not 0 (the "Finish picking commits" option).
					pickedIndexes = append(pickedIndexes, pickedIndex)
				}
			}

			pickedIndexes = helpers.GetUniqueIntegers(pickedIndexes)
		}
	case MigrateModeSingle:
		{
			// Select a single commit.
			prompt := promptui.Select{
				Label: "Select a commit to migrate",
				Items: gitLogArray,
				Stdin: reader,
			}

			pickedIndex, _, err := prompt.Run()
			if err != nil {
				panic(err)
			}

			pickedIndexes = append(pickedIndexes, pickedIndex)
		}
	case MigrateModeRange:
		{
			// Select the start commit.
			prompt := promptui.Select{
				Label: "Select the start commit",
				Items: gitLogArray,
				Stdin: reader,
			}

			startCommitIndex, _, err := prompt.Run()
			if err != nil {
				panic(err)
			}

			// Select the end commit.
			prompt = promptui.Select{
				Label: "Select the end commit",
				Items: gitLogArray,
				Stdin: reader,
			}

			endCommitIndex, _, err := prompt.Run()
			if err != nil {
				panic(err)
			}

			// Validate the picked options.
			if startCommitIndex <= endCommitIndex {
				panic(errors.New("start commit should be older than the end commit"))
			}

			pickedIndexes = helpers.GetRangeArrayFromTwoIntegers(startCommitIndex, endCommitIndex)
		}
	}

	// Sort the indexes from old to new (the bigger the index, the older they are).
	sort.Slice(pickedIndexes, func(i, j int) bool {
		return pickedIndexes[i] > pickedIndexes[j]
	})

	// Extract commits from the picked indexes.
	return helpers.PickRevisionsFromCommits(gitLogArray, pickedIndexes)
}

// Helper functions.
var migrateModeMap = map[string]MigrateModeType{
	"single":   MigrateModeSingle,
	"multiple": MigrateModeMultiple,
	"range":    MigrateModeRange,
}

func parseMigrateMode(mode string) (MigrateModeType, error) {
	result, ok := migrateModeMap[mode]
	if !ok {
		return -1, errors.New("not found")
	}

	return result, nil
}
