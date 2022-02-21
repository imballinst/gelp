package gelp

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"

	"github.com/manifoldco/promptui"
)

var gitLogArray = []string{
	"099e593: feature: add gelp squashto (#2) (2022-02-20T11:42:57+07:00)",
	"3b16259: checkpoint for squash new (2022-02-19T19:28:03+07:00)",
	"f38d9ee: remove unused (2022-02-19T19:05:00+07:00)",
	"asd123: (no commit title) (2022-02-19T18:05:00+07:00)",
}

func TestResolveRevisionsFromMigratePromptSingle(t *testing.T) {
	b := bytes.NewBuffer([]byte(string(promptui.KeyNext) + string(promptui.KeyEnter)))
	reader := ioutil.NopCloser(b)

	revisions, err := ResolveRevisionsFromMigratePrompt(gitLogArray, MigrateModeSingle, []io.ReadCloser{reader})
	if err != nil {
		t.Log("expected to not output any error")
		t.Fail()
	}

	if len(revisions) > 1 {
		t.Log(fmt.Sprintf("Expected %d, got %d instead", 1, len(revisions)))
		t.Fail()
	}

	if revisions[0] != "3b16259" {
		t.Log(fmt.Sprintf("Expected %s, got %s instead", "3b16259", revisions[0]))
		t.Fail()
	}
}

func TestResolveRevisionsFromMigratePromptMultiple(t *testing.T) {
	// Select 1st commit.
	b := bytes.NewBuffer([]byte(string(promptui.KeyNext) + string(promptui.KeyEnter)))
	reader := ioutil.NopCloser(b)

	// Select 4th commit.
	b = bytes.NewBuffer([]byte(
		string(promptui.KeyNext) +
			string(promptui.KeyNext) +
			string(promptui.KeyNext) +
			string(promptui.KeyNext) +
			string(promptui.KeyEnter),
	))
	reader2 := ioutil.NopCloser(b)

	// Select 3rd commit.
	b = bytes.NewBuffer([]byte(
		string(promptui.KeyNext) +
			string(promptui.KeyNext) +
			string(promptui.KeyNext) +
			string(promptui.KeyEnter),
	))
	reader3 := ioutil.NopCloser(b)

	// Finish selection by selecting the first commit.
	b = bytes.NewBuffer([]byte(string(promptui.KeyEnter)))
	reader4 := ioutil.NopCloser(b)

	revisions, err := ResolveRevisionsFromMigratePrompt(gitLogArray, MigrateModeMultiple, []io.ReadCloser{
		reader,
		reader2,
		reader3,
		reader4,
	})
	if err != nil {
		t.Log("expected to not output any error")
		t.Fail()
	}

	// Compare.
	exp := []string{
		"asd123",
		"f38d9ee",
		"099e593",
	}
	if len(revisions) != 3 {
		t.Log(fmt.Sprintf("Expected %d, got %d instead", len(exp), len(revisions)))
		t.Fail()
	}

	for i := range exp {
		if exp[i] != revisions[i] {
			t.Log(fmt.Sprintf("Expected %s, got %s instead", exp[i], revisions[i]))
			t.Fail()
		}
	}
}

func TestResolveRevisionsFromMigratePromptRange(t *testing.T) {
	// Select 2nd commit.
	b := bytes.NewBuffer([]byte(string(promptui.KeyNext) + string(promptui.KeyEnter)))
	reader := ioutil.NopCloser(b)

	// Select 4th commit.
	b = bytes.NewBuffer([]byte(
		string(promptui.KeyNext) +
			string(promptui.KeyNext) +
			string(promptui.KeyNext) +
			string(promptui.KeyEnter),
	))
	reader2 := ioutil.NopCloser(b)

	revisions, err := ResolveRevisionsFromMigratePrompt(gitLogArray, MigrateModeRange, []io.ReadCloser{
		reader,
		reader2,
	})
	if err != nil {
		t.Log("expected to not output any error")
		t.Fail()
	}

	// Compare.
	exp := []string{
		"asd123",
		"f38d9ee",
		"3b16259",
	}
	if len(revisions) != 3 {
		t.Log(fmt.Sprintf("Expected %d, got %d instead", len(exp), len(revisions)))
		t.Fail()
	}

	for i := range exp {
		if exp[i] != revisions[i] {
			t.Log(fmt.Sprintf("Expected %s, got %s instead", exp[i], revisions[i]))
			t.Fail()
		}
	}
}
