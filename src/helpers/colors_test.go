// This file is only used to test colors, and
// not really about testing the functions.
package gelp

import (
	"fmt"
	"testing"
)

func TestBlueColor(t *testing.T) {
	fmt.Println("Test blue color", GetBlueText("hello"), "After blue color")
}

func TestRedColor(t *testing.T) {
	fmt.Println("Test red color", GetRedText("hello"), "After red color")
}

func TestGreenColor(t *testing.T) {
	fmt.Println("Test green color", GetGreenText("hello"), "After green color")
}
