// This file should be tested in the end-to-end test.
// Testing this in unit test is a pain.
package gelp

import (
	"fmt"
	"strings"
)

func Migrate(targetBranch string, baseBranch string, pickedCommit string) error {
	commitArray := strings.SplitN(pickedCommit, " ", 2)
	commitRevision := commitArray[0]

	fmt.Println("git migrate:", "Getting current branch...")
	currentBranchOutput, err := ExecCommand("git rev-parse --abbrev-ref HEAD")
	if err != nil {
		return err
	}

	// Check if the branch exists.
	fmt.Println("git migrate:", (fmt.Sprintf("git rev-parse --quiet --verify %s", targetBranch)))
	_, err = ExecCommand(fmt.Sprintf("git rev-parse --quiet --verify %s", targetBranch))

	if err != nil {
		// Create a new branch, if the target branch doesn't exist.
		fmt.Println("git migrate: Branch doesn't exist, creating one...")
		fmt.Println("git migrate:", (fmt.Sprintf("git checkout -b %s %s", targetBranch, baseBranch)))
		_, err = ExecCommand(fmt.Sprintf("git checkout -b %s %s", targetBranch, baseBranch))
		if err != nil {
			return err
		}
	} else {
		// Branch exists.
		_, err = ExecCommand(fmt.Sprintf("git checkout %s", targetBranch))
		if err != nil {
			return err
		}
	}

	// Cherry-pick the commit.
	fmt.Println("git migrate:", (fmt.Sprintf("git cherry-pick %s", commitRevision)))
	_, err = ExecCommand(fmt.Sprintf("git cherry-pick %s", commitRevision))
	if err != nil {
		return err
	}

	// Go back to the old branch.
	fmt.Println("git migrate:", (fmt.Sprintf("git switch %s", currentBranchOutput)))
	_, err = ExecCommand(fmt.Sprintf("git switch %s", currentBranchOutput))
	if err != nil {
		return err
	}

	return nil
}
