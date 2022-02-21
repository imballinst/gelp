package helpers

import (
	"fmt"
	"testing"
)

func TestGetUniqueIntegers(t *testing.T) {
	integers := []int{
		6, 1, 2, 3, 3, 4, 4, 1, 2, 5, 6, 6,
	}

	exp := []int{
		6, 1, 2, 3, 4, 5,
	}
	actual := GetUniqueIntegers(integers)

	for i := range exp {
		if exp[i] != actual[i] {
			t.Log(fmt.Sprintf("Expected %d, got %d instead", exp[i], actual[i]))
			t.Fail()
		}
	}
}

func TestGetRangeArrayFromTwoIntegers(t *testing.T) {
	// Test ascending.
	exp := []int{5, 4, 3, 2}
	actual := GetRangeArrayFromTwoIntegers(5, 2)
	for i := range exp {
		if exp[i] != actual[i] {
			t.Log(fmt.Sprintf("Expected %d, got %d instead", exp[i], actual[i]))
			t.Fail()
		}
	}

	// Test descending.
	exp = []int{8, 7, 6, 5, 4}
	actual = GetRangeArrayFromTwoIntegers(8, 4)
	for i := range exp {
		if exp[i] != actual[i] {
			t.Log(fmt.Sprintf("Expected %d, got %d instead", exp[i], actual[i]))
			t.Fail()
		}
	}
}
