package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestEqual(t *testing.T) {
	testCases := map[string]map[string]error{
		"abc": {"abc": nil, "def": ErrNotEqual.With("abc"), "xyz": ErrNotEqual.With("abc")},
		"def": {"abc": ErrNotEqual.With("def"), "def": nil, "xyz": ErrNotEqual.With("def")},
		"xyz": {"abc": ErrNotEqual.With("xyz"), "def": ErrNotEqual.With("xyz"), "xyz": nil},
	}

	for setup, values := range testCases {
		testEqual := Equal(setup)

		for input, want := range values {
			t.Run(fmt.Sprintf("%s/%s", setup, input), func(t *testing.T) {
				got := testEqual(input)

				if !errors.Is(got, want) {
					t.Error("got", got)
					t.Error("want", want)
				}
			})
		}
	}
}
