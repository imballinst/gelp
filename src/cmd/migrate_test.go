package gelp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/manifoldco/promptui"
)

type ClosingBuffer struct {
	*bytes.Buffer
}

func (cb ClosingBuffer) Close() error {
	return nil
}

var gitLogArray = []string{
	"099e593: feature: add gelp squashto (#2) (2022-02-20T11:42:57+07:00)",
	"3b16259: checkpoint for squash new (2022-02-19T19:28:03+07:00)",
	"f38d9ee: remove unused (2022-02-19T19:05:00+07:00)",
	"asd123: (no commit title) (2022-02-19T18:05:00+07:00)",
}

func TestResolveRevisionsFromMigratePromptSingle(t *testing.T) {
	b := bytes.NewBuffer([]byte(string(promptui.KeyNext) + string(promptui.KeyEnter)))
	reader := ioutil.NopCloser(
		b,
	)

	revisions := ResolveRevisionsFromMigratePrompt(gitLogArray, MigrateModeSingle, reader)
	if len(revisions) > 1 {
		t.Log(fmt.Sprintf("Expected %d, got %d instead", 1, len(revisions)))
		t.Fail()
	}

	if revisions[0] != "3b16259" {
		t.Log(fmt.Sprintf("Expected %s, got %s instead", "3b16259", revisions[0]))
		t.Fail()
	}
}
