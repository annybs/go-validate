package validate

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleUUID() {
	fmt.Println(UUID("not a uuid"))
	// Output: invalid UUID
}

func TestUUID(t *testing.T) {
	testCases := map[string]error{
		"00000000-0000-0000-0000-000000000000": nil,
		"01234567-89ab-cdef-0123-456789abcdef": nil,
		"abcdef01-2345-6789-abcd-ef0123456789": nil,

		"not a uuid":                           ErrInvalidUUID,
		"00000000-00-0000-0000-00000000000000": ErrInvalidUUID,
		"00000000000000000000000000000000":     ErrInvalidUUID,
		"01234567-89ab-cdef-ghij-klmnopqrstuv": ErrInvalidUUID,
	}

	for input, want := range testCases {
		t.Run(input, func(t *testing.T) {
			got := UUID(input)

			if !errors.Is(got, want) {
				t.Error("got", got)
				t.Error("want", want)
			}
		})
	}
}
