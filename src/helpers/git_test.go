package gelp

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func init() {
	if os.Getenv("GO_WANT_HELPER_PROCESS") == "1" {
		return
	}
}

type GitTestRunner struct{}

func (r GitTestRunner) Migrate(targetBranch string, baseBranch string, pickedCommit string) (string, error) {
	cs := []string{"-test.run=TestHelperProcess", "--"}
	cs = append(cs, targetBranch, pickedCommit, "--base", baseBranch)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func TestMigrate(t *testing.T) {
	gitRunner = GitTestRunner{}
	output, err := gitRunner.Migrate("hotfix", "samplecommit", "dev")
	fmt.Println(output, err)
}

// Source: https://github.com/golang/go/blob/master/src/os/exec/exec_test.go#L681.
// And: 	https://joshrendek.com/2014/06/go-lang-mocking-exec-dot-command-using-interfaces/.
func TestHelperProcess(*testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	defer os.Exit(0)

	args := os.Args
	cmd, args := args[0], args[1:]

	iargs := ""
	for _, s := range args {
		iargs = iargs + s
	}
	fmt.Println(cmd, iargs)
}
