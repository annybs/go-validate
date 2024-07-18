package validate

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleEmail() {
	fmt.Println(Email("not an email"))
	// Output: invalid email address
}

func FuzzEmail(f *testing.F) {
	want := ErrInvalidEmail

	f.Fuzz(func(t *testing.T, input string) {
		got := Email(input)

		if !errors.Is(got, want) {
			t.Error("got", got)
			t.Error("want", want)
		}
	})
}

func TestEmail(t *testing.T) {
	testCases := map[string]error{
		"test@example.com":                      nil,
		"firstname.lastname@some-website.co.uk": nil,

		"not an email":    ErrInvalidEmail,
		"testexample.com": ErrInvalidEmail,
	}

	for input, want := range testCases {
		t.Run(input, func(t *testing.T) {
			got := Email(input)

			if !errors.Is(got, want) {
				t.Error("got", got)
				t.Error("want", want)
			}
		})
	}
}
