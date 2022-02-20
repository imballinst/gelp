// This file should be tested in the end-to-end test.
// Testing this in unit test is a pain.
package gelp

import (
	"fmt"
	"strings"
)

// Commands used in the CLI.
func Migrate(targetBranch string, baseBranch string, revisions []string) error {
	// Check/create and checkout to that branch.
	err := checkoutOtherBranchOrCreateNew("migrate", targetBranch, baseBranch)
	if err != nil {
		return err
	}

	cherrypickCommand := "git cherry-pick " + strings.Join(revisions, " ")

	// Cherry-pick the commits.
	_, err = doAndLog("migrate", cherrypickCommand)
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

// Semi-helper functions. Used to create arguments passed to functions above.
// These functions are exported so that we can compose the functions better.
//
// We are using this command: "git log --date=iso-strict --pretty='format:%cd %h %s'".
// Which outputs (for example):
// 2022-02-20T11:42:57+07:00 099e593 feature: add gelp squashto (#2)
// 2022-02-19T19:28:03+07:00 3b16259 checkpoint for squash new
// 2022-02-19T19:05:00+07:00 f38d9ee remove unused
func ExtractRevisionAndTitleFromCommits(commits []string, isWithDate bool) []string {
	var result []string

	for _, commit := range commits {
		revisionAndMessage := strings.SplitN(commit, " ", 3)
		var entry string
		var commitTitle string

		// Check the length of the split string.
		// There can be a chance where the commit title is empty.
		if len(revisionAndMessage) == 2 {
			// Has only 2 segments (the commit title is empty).
			commitTitle = "(no commit title)"
		} else {
			// Has 3 segments.
			commitTitle = revisionAndMessage[2]
		}

		// Depending on the `isWithDate` flag, change entry format.
		if isWithDate {
			entry = fmt.Sprintf("%s: %s (%s)", revisionAndMessage[1], commitTitle, revisionAndMessage[0])
		} else {
			entry = fmt.Sprintf("%s: %s", revisionAndMessage[1], commitTitle)
		}

		result = append(result, entry)
	}

	return result
}

func PickRevisionsFromCommits(commits []string, indexes []int) []string {
	var revisions []string

	for _, index := range indexes {
		revisions = append(revisions, commits[index])
	}

	return revisions
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
