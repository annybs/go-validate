package validate

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleIn() {
	testIn := In("abc", "def", "xyz")
	fmt.Println(testIn("123"))
	// Output: not allowed
}

func TestIn(t *testing.T) {
	testIn := In("abc", "def", "xyz")

	testCases := map[string]error{
		"abc": nil,
		"def": nil,
		"xyz": nil,

		"abcd": ErrValueNotAllowed,
		"123":  ErrValueNotAllowed,
		"":     ErrValueNotAllowed,
	}

	for input, want := range testCases {
		t.Run(input, func(t *testing.T) {
			got := testIn(input)

			if !errors.Is(got, want) {
				t.Error("got", got)
				t.Error("want", want)
			}
		})
	}
}
