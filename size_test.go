package validate

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleMaxSize() {
	testMaxSize := MaxSize[string](3)
	fmt.Println(testMaxSize([]string{"abc", "def", "ghi", "jkl"}))
	// Output: must have no more than 3 items
}

func ExampleMinSize() {
	testMaxSize := MinSize[string](3)
	fmt.Println(testMaxSize([]string{"abc", "def"}))
	// Output: must have at least 3 items
}

func TestMaxSize(t *testing.T) {
	testCases := map[int]map[int]error{
		0:   {0: nil, 1: ErrMustHaveFewerItems.With(0), 2: ErrMustHaveFewerItems.With(0), 10: ErrMustHaveFewerItems.With(0)},
		1:   {0: nil, 1: nil, 2: ErrMustHaveFewerItems.With(1), 10: ErrMustHaveFewerItems.With(1)},
		2:   {0: nil, 1: nil, 2: nil, 10: ErrMustHaveFewerItems.With(2)},
		10:  {0: nil, 1: nil, 2: nil, 10: nil},
		100: {0: nil, 1: nil, 2: nil, 10: nil},
	}

	for max, values := range testCases {
		testMaxSize := MaxSize[int](max)

		for l, want := range values {
			t.Run(fmt.Sprintf("%d/%d", max, l), func(t *testing.T) {
				input := []int{}
				for i := 0; i < l; i++ {
					input = append(input, i)
				}

				got := testMaxSize(input)

				if !errors.Is(got, want) {
					t.Error("got", got)
					t.Error("want", want)
				}
			})
		}
	}
}

func TestMinSize(t *testing.T) {
	testCases := map[int]map[int]error{
		0:   {0: nil, 1: nil, 2: nil, 10: nil},
		1:   {0: ErrMustHaveMoreItems.With(1), 1: nil, 2: nil, 10: nil},
		2:   {0: ErrMustHaveMoreItems.With(2), 1: ErrMustHaveMoreItems.With(2), 2: nil, 10: nil},
		10:  {0: ErrMustHaveMoreItems.With(10), 1: ErrMustHaveMoreItems.With(10), 2: ErrMustHaveMoreItems.With(10), 10: nil},
		100: {0: ErrMustHaveMoreItems.With(100), 1: ErrMustHaveMoreItems.With(100), 2: ErrMustHaveMoreItems.With(100), 10: ErrMustHaveMoreItems.With(100)},
	}

	for min, values := range testCases {
		testMinSize := MinSize[int](min)

		for l, want := range values {
			t.Run(fmt.Sprintf("%d/%d", min, l), func(t *testing.T) {
				input := []int{}
				for i := 0; i < l; i++ {
					input = append(input, i)
				}

				got := testMinSize(input)

				if !errors.Is(got, want) {
					t.Error("got", got)
					t.Error("want", want)
				}
			})
		}
	}
}
