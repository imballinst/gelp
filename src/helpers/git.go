// This file should be tested in the end-to-end test.
// Testing this in unit test is a pain.
package helpers

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type GitLog struct {
	Date   string
	Commit string
	Title  string
	Branch string
}

// Commands used in the CLI.
func Migrate(targetBranch string, baseBranch string, revisions []string) error {
	// Check/create and checkout to that branch.
	err := checkoutOtherBranchOrCreateNew("migrate", targetBranch, baseBranch)
	if err != nil {
		return err
	}

	cherrypickCommand := "git cherry-pick " + strings.Join(revisions, " ")

	// Cherry-pick the commits.
	_, err = DoAndLog("migrate", cherrypickCommand)
	if err != nil {
		return err
	}

	return nil
}

func Squashto(targetBranch string, baseBranch string) error {
	currentBranchOutput, err := getCurrentBranch("squashto")
	if err != nil {
		return err
	}

	// Check/create and checkout to that branch.
	err = checkoutOtherBranchOrCreateNew("squashto", targetBranch, baseBranch)
	if err != nil {
		return err
	}

	// Squash merge.
	_, err = DoAndLog("squashto", fmt.Sprintf("git merge --squash %s", currentBranchOutput))
	if err != nil {
		return err
	}

	return nil
}

func Postmerge(remote, baseBranch string) error {
	currentBranchOutput, err := getCurrentBranch("postmerge")
	if err != nil {
		return err
	}

	// Checkout to base branch.
	err = CheckoutBranch("postmerge", baseBranch)
	if err != nil {
		return err
	}

	// Pull the changes.
	_, err = DoAndLog("postmerge", fmt.Sprintf("git pull %s %s", remote, baseBranch))
	if err != nil {
		return err
	}

	// Delete the old branch.
	_, err = DoAndLog("postmerge", fmt.Sprintf("git branch -D %s", currentBranchOutput))
	if err != nil {
		return err
	}

	return nil
}

func Fresh(targetBranch, baseBranch string) error {
	return checkoutNewBranch("fresh", targetBranch, baseBranch)
}

func UpdateBranch(remote, targetBranch string) error {
	currentBranchOutput, err := getCurrentBranch("updatebranch")
	if err != nil {
		return err
	}

	if currentBranchOutput == targetBranch || targetBranch == "" {
		// When target branch is the same or is empty, then we update the current branch instead.
		_, err = DoAndLog("updatebranch", fmt.Sprintf("git pull %s %s", remote, currentBranchOutput))
		return err
	}

	_, err = DoAndLog("updatebranch", fmt.Sprintf("git fetch %s %s:%s", remote, targetBranch, targetBranch))
	return err
}

func BranchList() ([]GitLog, error) {
	branchListOutput, err := DoAndLog("branch", "git branch")
	if err != nil {
		return nil, err
	}

	branchListOutput = strings.Replace(branchListOutput, "*", "", -1)

	var gitLogInfos = []GitLog{}
	branchListArray := strings.Split(strings.Trim(branchListOutput, "\n"), "\n")

	for _, branch := range branchListArray {
		log, err := GetCommitsFromSource(strings.Trim(branch, " "), 1)
		if err != nil {
			return nil, err
		}
		gitLogInfos = append(gitLogInfos, log)
	}

	sort.Slice(gitLogInfos, func(i, j int) bool {
		parsed1, err := time.Parse(DATE_FORMAT, gitLogInfos[i].Date)
		if err != nil {
			return false
		}
		parsed2, err := time.Parse(DATE_FORMAT, gitLogInfos[j].Date)
		if err != nil {
			return false
		}

		return parsed1.UnixMicro() > parsed2.UnixMicro()
	})

	return gitLogInfos, nil
}

// Semi-helper functions. Used to create arguments passed to functions above.
// These functions are exported so that we can compose the functions better.
//
// We are using this command: "git log --date=iso-strict --pretty='format:%cd %h %s'".
// Which outputs (for example):
// 2022-02-20T11:42:57+07:00 099e593 feature: add gelp squashto (#2)
// 2022-02-19T19:28:03+07:00 3b16259 checkpoint for squash new
// 2022-02-19T19:05:00+07:00 f38d9ee remove unused
func ExtractRevisionAndTitleFromCommits(commits []GitLog, isWithDate bool) []string {
	var result []string

	for _, commit := range commits {
		var entry string
		var commitTitle string

		// Check the length of the split string.
		// There can be a chance where the commit title is empty.
		if commit.Title == "" {
			// Has only 2 segments (the commit title is empty).
			commitTitle = "(no commit title)"
		} else {
			// Has 3 segments.
			commitTitle = commit.Title
		}

		// Depending on the `isWithDate` flag, change entry format.
		if isWithDate {
			entry = fmt.Sprintf("%s: %s (%s)", commit.Commit, commitTitle, commit.Date)
		} else {
			entry = fmt.Sprintf("%s: %s", commit.Commit, commitTitle)
		}

		result = append(result, entry)
	}

	return result
}

// Pick only the revisions from this format (for example):
// 099e593: feature: add gelp squashto (#2) (2022-02-20T11:42:57+07:00)
// 3b16259: checkpoint for squash new (2022-02-19T19:28:03+07:00)
// f38d9ee: remove unused (2022-02-19T19:05:00+07:00)"
func PickRevisionsFromCommits(commits []string, indexes []int) []string {
	var revisions []string

	for _, index := range indexes {
		commitSplitArray := strings.SplitN(commits[index], ": ", 2)
		revisions = append(revisions, commitSplitArray[0])
	}

	return revisions
}

func GetCommitsFromSource(source string, num int) (GitLog, error) {
	command := "git log --date=iso-strict --pretty='%cd %h %s'"
	if source != "" {
		command = fmt.Sprintf("%s --source %s", command, source)
	}

	if num != -1 {
		command = fmt.Sprintf("%s -n %d", command, num)
	}

	var gitLogInfo = GitLog{}

	gitLog, err := ExecCommand(command)
	if err != nil {
		return gitLogInfo, err
	}

	spl := strings.SplitN(gitLog, " ", 3)
	gitLogInfo = GitLog{
		Date:   spl[0],
		Commit: spl[1],
		Title:  spl[2],
		Branch: source,
	}

	return gitLogInfo, nil
}

func CheckoutBranch(label, targetBranch string) error {
	_, err := DoAndLog(label, fmt.Sprintf("git checkout %s", targetBranch))
	if err != nil {
		return err
	}

	return nil
}

// Helper functions.
func checkoutNewBranch(label, targetBranch, baseBranch string) error {
	_, err := DoAndLog(label, fmt.Sprintf("git checkout -b %s %s", targetBranch, baseBranch))
	if err != nil {
		return err
	}

	return nil
}

func checkoutOtherBranchOrCreateNew(label string, targetBranch string, baseBranch string) error {
	// Check if the branch exists.
	verifyBranchExistsCommand := fmt.Sprintf("git rev-parse --quiet --verify %s", targetBranch)
	_, err := DoAndLog(label, verifyBranchExistsCommand)

	if err != nil {
		Log(label, "Branch doesn't exist, creating one...")

		// Create a new branch, if the target branch doesn't exist.
		err = checkoutNewBranch(label, targetBranch, baseBranch)
		if err != nil {
			return err
		}
	} else {
		// Branch exists.
		err = CheckoutBranch(label, targetBranch)
		if err != nil {
			return err
		}
	}

	return nil
}

func getCurrentBranch(label string) (string, error) {
	currentBranchOutput, err := DoAndLog(label, "git branch --show-current")
	if err != nil {
		return "", err
	}

	return currentBranchOutput, nil
}
