package validate

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleChars() {
	testChars := Chars("0123456789abcdef")
	fmt.Println(testChars("invalid input"))
	// Output: contains disallowed characters
}

func ExampleExceptChars() {
	testExceptChars := Chars("0123456789abcdef")
	fmt.Println(testExceptChars("invalid input"))
	// Output: contains disallowed characters
}

func TestChars(t *testing.T) {
	testCases := map[string]map[string]error{
		"0123456789abcdef": {"abcd1234": nil, "abcd 1234": ErrDisallowedChars, "ghijklmno": ErrDisallowedChars},
	}

	for setup, values := range testCases {
		testChars := Chars(setup)

		for input, want := range values {
			t.Run(input, func(t *testing.T) {
				got := testChars(input)

				if !errors.Is(got, want) {
					t.Error("got", got)
					t.Error("want", want)
				}
			})
		}
	}
}

func TestExceptChars(t *testing.T) {
	testCases := map[string]map[string]error{
		"0123456789abcdef": {"abcd1234": ErrDisallowedChars, "abcd 1234": ErrDisallowedChars, "ghijklmno": nil},
	}

	for setup, values := range testCases {
		testExceptChars := ExceptChars(setup)

		for input, want := range values {
			t.Run(input, func(t *testing.T) {
				got := testExceptChars(input)

				if !errors.Is(got, want) {
					t.Error("got", got)
					t.Error("want", want)
				}
			})
		}
	}
}
