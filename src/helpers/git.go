// This file should be tested in the end-to-end test.
// Testing this in unit test is a pain.
package gelp

import (
	"fmt"
	"strings"
)

func Migrate(targetBranch string, baseBranch string, startCommit string, endCommit string) error {
	startCommitArray := strings.SplitN(startCommit, " ", 2)
	startCommitRevision := startCommitArray[0]
	endCommitArray := strings.SplitN(endCommit, " ", 2)
	endCommitRevision := endCommitArray[0]

	// Check/create and checkout to that branch.
	err := checkoutOtherBranchOrCreateNew("migrate", targetBranch, baseBranch)
	if err != nil {
		return err
	}

	// Cherry-pick the commits.
	_, err = doAndLog("migrate", fmt.Sprintf("git cherry-pick %s^..%s", startCommitRevision, endCommitRevision))
	if err != nil {
		return err
	}

	return nil
}

func Squashto(targetBranch string, baseBranch string) error {
	currentBranchOutput, err := doAndLog("squashto", "git rev-parse --abbrev-ref HEAD")
	if err != nil {
		return err
	}

	// Check/create and checkout to that branch.
	err = checkoutOtherBranchOrCreateNew("squashto", targetBranch, baseBranch)
	if err != nil {
		return err
	}

	// Squash merge.
	_, err = doAndLog("squashto", fmt.Sprintf("git merge --squash %s", currentBranchOutput))
	if err != nil {
		return err
	}

	return nil
}

// Helper functions.
func doAndLog(label string, command string) (string, error) {
	log(label, command)
	return ExecCommand(command)
}

func log(label string, text string) {
	fmt.Printf("gelp %s: %s\n", label, text)
}

func checkoutOtherBranchOrCreateNew(label string, targetBranch string, baseBranch string) error {
	// Check if the branch exists.
	verifyBranchExistsCommand := fmt.Sprintf("git rev-parse --quiet --verify %s", targetBranch)
	_, err := doAndLog(label, verifyBranchExistsCommand)

	if err != nil {
		log(label, "Branch doesn't exist, creating one...")

		// Create a new branch, if the target branch doesn't exist.
		createNewBranchCommand := fmt.Sprintf("git checkout -b %s %s", targetBranch, baseBranch)
		_, err = doAndLog(label, createNewBranchCommand)
		if err != nil {
			return err
		}
	} else {
		// Branch exists.
		_, err = doAndLog(label, fmt.Sprintf("git checkout %s", targetBranch))
		if err != nil {
			return err
		}
	}

	return nil
}
