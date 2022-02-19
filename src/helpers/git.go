package gelp

import "fmt"

func Migrate(targetBranch string, baseBranch string, pickedCommit string) error {
	currentBranch, err := ExecCommand("git rev-parse --abbrev-ref HEAD")
	if err != nil {
		return err
	}

	// Check if the branch exists.
	output, err := ExecCommand(fmt.Sprintf("git rev-parse --quiet --verify %s", targetBranch))
	if err != nil {
		return err
	}

	// Create a new branch, if the target branch doesn't exist.
	if output == "" {
		_, err = ExecCommand(fmt.Sprintf("git checkout -b %s %s", targetBranch, baseBranch))
		if err != nil {
			return err
		}
	}

	// Cherry-pick the commit.
	_, err = ExecCommand(fmt.Sprintf("git cherry-pick %s", pickedCommit))
	if err != nil {
		return err
	}

	// Go back to the old branch.
	_, err = ExecCommand(fmt.Sprintf("git switch %s", currentBranch))
	if err != nil {
		return err
	}

	return nil
}
