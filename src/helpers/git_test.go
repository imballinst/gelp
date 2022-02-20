package gelp

import (
	"fmt"
	"testing"
)

// We are only testing "non-CLI binary" functions because
// they will be tested in the end-to-end test instead.
func TestExtractRevisionAndTitleFromCommits(t *testing.T) {
	// With date.
	commits := []string{
		"2022-02-20T11:42:57+07:00 099e593 feature: add gelp squashto (#2)",
		"2022-02-19T19:28:03+07:00 3b16259 checkpoint for squash new",
		"2022-02-19T19:05:00+07:00 f38d9ee remove unused",
		"2022-02-19T18:05:00+07:00 asd123", // Empty commit message, intentionally tested.
	}

	exp := []string{
		"099e593: feature: add gelp squashto (#2) (2022-02-20T11:42:57+07:00)",
		"3b16259: checkpoint for squash new (2022-02-19T19:28:03+07:00)",
		"f38d9ee: remove unused (2022-02-19T19:05:00+07:00)",
		"asd123: (no commit title) (2022-02-19T18:05:00+07:00)",
	}
	actual := ExtractRevisionAndTitleFromCommits(commits, true)

	for i := range exp {
		if exp[i] != actual[i] {
			t.Log(fmt.Sprintf("Expected %s, got %s instead", exp[i], actual[i]))
			t.Fail()
		}
	}

	// Without date.
	exp = []string{
		"099e593: feature: add gelp squashto (#2)",
		"3b16259: checkpoint for squash new",
		"f38d9ee: remove unused",
		"asd123: (no commit title)",
	}
	actual = ExtractRevisionAndTitleFromCommits(commits, false)

	for i := range exp {
		if exp[i] != actual[i] {
			t.Log(fmt.Sprintf("Expected %s, got %s instead", exp[i], actual[i]))
			t.Fail()
		}
	}
}

func TestPickRevisionsFromCommits(t *testing.T) {
	commits := []string{
		"2022-02-20T11:42:57+07:00 099e593 feature: add gelp squashto (#2)",
		"2022-02-19T19:28:03+07:00 3b16259 checkpoint for squash new",
		"2022-02-19T19:05:00+07:00 f38d9ee remove unused",
		"2022-02-19T18:05:00+07:00 asd123", // Empty commit message, intentionally tested.
	}
	indexes := []int{0, 1, 3}

	exp := []string{
		"2022-02-20T11:42:57+07:00 099e593 feature: add gelp squashto (#2)",
		"2022-02-19T19:28:03+07:00 3b16259 checkpoint for squash new",
		"2022-02-19T18:05:00+07:00 asd123", // Empty commit message, intentionally tested.
	}
	actual := PickRevisionsFromCommits(commits, indexes)

	for i := range exp {
		if exp[i] != actual[i] {
			t.Log(fmt.Sprintf("Expected %s, got %s instead", exp[i], actual[i]))
			t.Fail()
		}
	}
}
