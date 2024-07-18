package validate

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleAll() {
	testAll := All(MinLength(4), Chars("0123456789abcdef"))
	fmt.Println(testAll("invalid input"))
	// Output: contains disallowed characters
}

func TestAll(t *testing.T) {
	testAll := All(
		MinLength(4),
		MaxLength(8),
		Chars("0123456789abcdef"),
		In("abcd", "abcdef", "12345678"),
	)

	testCases := map[string]error{
		"abcd":     nil,
		"abcdef":   nil,
		"12345678": nil,

		"abc":       ErrMustBeLonger.With(4),
		"abcdef012": ErrMustBeShorter.With(8),
		"abcdefgh":  ErrDisallowedChars,
		"01abcd":    ErrValueNotAllowed,
	}

	for input, want := range testCases {
		t.Run(input, func(t *testing.T) {
			got := testAll(input)

			if !errors.Is(got, want) {
				t.Error("got", got)
				t.Error("want", want)
			}
		})
	}
}
