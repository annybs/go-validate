package validate

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleMaxLength() {
	testMaxLength := MaxLength(8)
	fmt.Println(testMaxLength("this string is too long"))
	// Output: must contain no more than 8 characters
}

func ExampleMinLength() {
	testMinLength := MinLength(8)
	fmt.Println(testMinLength("2short"))
	// Output: must contain at least 8 characters
}

func TestMaxLength(t *testing.T) {
	testCases := map[int]map[string]error{
		8: {"abcd": nil, "abcdefgh": nil, "abcd efg": nil, "abcdefghi": ErrMustBeShorter.With(8)},
	}

	for setup, values := range testCases {
		testMaxLength := MaxLength(setup)

		for input, want := range values {
			t.Run(fmt.Sprintf("%d/%s", setup, input), func(t *testing.T) {
				got := testMaxLength(input)

				if !errors.Is(got, want) {
					t.Error("got", got)
					t.Error("want", want)
				}
			})
		}
	}
}

func TestMinLength(t *testing.T) {
	testCases := map[int]map[string]error{
		8: {"abcd": ErrMustBeLonger.With(8), "abcdefgh": nil, "abcd efg": nil, "abcdefghi": nil},
	}

	for setup, values := range testCases {
		testMinLength := MinLength(setup)

		for input, want := range values {
			t.Run(fmt.Sprintf("%d/%s", setup, input), func(t *testing.T) {
				got := testMinLength(input)

				if !errors.Is(got, want) {
					t.Error("got", got)
					t.Error("want", want)
				}
			})
		}
	}
}
