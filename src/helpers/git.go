package gelp

import "fmt"

type GitRunner interface {
	Migrate(targetBranch string, baseBranch string, pickedCommit string) ([]byte, error)
}
type RealGitRunner struct{}

var gitRunner GitRunner

func (r RealGitRunner) Migrate(targetBranch string, baseBranch string, pickedCommit string) ([]byte, error) {
	currentBranch, err := ExecCommand("git rev-parse --abbrev-ref HEAD")
	if err != nil {
		return nil, err
	}

	// Check if the branch exists.
	output, err := ExecCommand(fmt.Sprintf("git rev-parse --quiet --verify %s", targetBranch))
	if err != nil {
		return nil, err
	}

	// Create a new branch, if the target branch doesn't exist.
	if output == "" {
		_, err = ExecCommand(fmt.Sprintf("git switch %s %s", targetBranch, baseBranch))
	}

	// Go to the target branch.

	// Cherry-pick the commit.
	_, err = ExecCommand()
}
