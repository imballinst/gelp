package gelp

import "fmt"

type GitRunner interface {
	Migrate(targetBranch string, baseBranch string, pickedCommit string) (string, error)
}

var gitRunner GitRunner

func Migrate(targetBranch string, baseBranch string, pickedCommit string) (string, error) {
	var output string

	currentBranchOutput, err := ExecCommand("git rev-parse --abbrev-ref HEAD")
	if err != nil {
		return output, err
	}

	output = fmt.Sprintln(output + currentBranchOutput)
	// Check if the branch exists.
	revisionOutput, err := ExecCommand(fmt.Sprintf("git rev-parse --quiet --verify %s", targetBranch))
	if err != nil {
		return output, err
	}

	output = fmt.Sprintln(output + revisionOutput)
	// Create a new branch, if the target branch doesn't exist.
	if output == "" {
		checkoutOutput, err := ExecCommand(fmt.Sprintf("git checkout -b %s %s", targetBranch, baseBranch))
		if err != nil {
			return output, err
		}

		output = fmt.Sprintln(output + checkoutOutput)
	}

	// Cherry-pick the commit.
	cherrypickOutput, err := ExecCommand(fmt.Sprintf("git cherry-pick %s", pickedCommit))
	if err != nil {
		return output, err
	}

	output = fmt.Sprintln(output + cherrypickOutput)
	// Go back to the old branch.
	oldBranchOutput, err := ExecCommand(fmt.Sprintf("git switch %s", currentBranchOutput))
	if err != nil {
		return output, err
	}

	output = fmt.Sprintln(output + oldBranchOutput)
	return output, nil
}
