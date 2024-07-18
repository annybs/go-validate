package validate

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleURL() {
	fmt.Println(URL("not a url"))
	// Output: invalid URL
}

func TestURL(t *testing.T) {
	testCases := map[string]error{
		"http://example.com":                    nil,
		"http://subdomain.example.com":          nil,
		"http://www.example.com/some-page.html": nil,

		"not a url":     ErrInvalidURL,
		"subdomain.com": ErrInvalidURL,
	}

	for input, want := range testCases {
		t.Run(input, func(t *testing.T) {
			got := URL(input)

			if !errors.Is(got, want) {
				t.Error("got", got)
				t.Error("want", want)
			}
		})
	}
}
