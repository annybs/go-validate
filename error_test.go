package validate

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrorIs(t *testing.T) {
	type TestCase struct {
		A    error
		B    error
		Want bool
	}

	testCases := []TestCase{
		// Want any validation error
		{A: Err, B: Err, Want: true},
		{A: ErrDisallowedChars, B: Err, Want: true},
		{A: ErrMustBeGreater, B: Err, Want: true},

		// Want specific validation error
		{A: ErrDisallowedChars, B: ErrDisallowedChars, Want: true},
		{A: ErrMustBeGreater, B: ErrMustBeGreater, Want: true},

		// Want not specific validation error
		{A: Err, B: ErrDisallowedChars},
		{A: Err, B: ErrMustBeGreater},
		{A: ErrMustBeGreater, B: ErrDisallowedChars},
		{A: ErrDisallowedChars, B: ErrMustBeGreater},

		// Want not any other error
		{A: ErrDisallowedChars, B: errors.New("contains disallowed characters")},
		{A: ErrMustBeGreater, B: errors.New("must be greater than %v")},
	}

	for _, testCase := range testCases {
		a, b, want := testCase.A, testCase.B, testCase.Want

		t.Run(fmt.Sprintf("%v/%v", a, b), func(t *testing.T) {
			got := errors.Is(a, b)

			if got != want {
				t.Error("got", got)
				t.Error("want", want)
			}
		})
	}
}
