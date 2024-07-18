package validate

import (
	"errors"
	"fmt"
	"testing"
)

func ExampleMax() {
	testMax := Max(10, true)
	fmt.Println(testMax(10))
	// Output: must be less than 10
}

func ExampleMin() {
	testMin := Min(10, false)
	fmt.Println(testMin(5))
	// Output: must be greater than or equal to 10
}

func TestMax(t *testing.T) {
	testCases := map[int]map[bool]map[int]error{
		10: {
			true:  {0: nil, 1: nil, 2: nil, 10: ErrMustBeLess.With(10), 100: ErrMustBeLess.With(10)},
			false: {0: nil, 1: nil, 2: nil, 10: nil, 100: ErrMustBeLessOrEqual.With(10)},
		},
	}

	for setup, subSetup := range testCases {
		for excl, values := range subSetup {
			testMax := Max(setup, excl)

			for input, want := range values {
				t.Run(fmt.Sprintf("%d/%v/%d", setup, excl, input), func(t *testing.T) {
					got := testMax(input)

					if !errors.Is(got, want) {
						t.Error("got", got)
						t.Error("want", want)
					}
				})
			}
		}
	}
}

func TestMaxFloat32(t *testing.T) {
	testCases := map[float32]map[bool]map[float32]error{
		10: {
			true:  {0: nil, 1: nil, 2: nil, 10: ErrMustBeLess.With(10), 100: ErrMustBeLess.With(10)},
			false: {0: nil, 1: nil, 2: nil, 10: nil, 100: ErrMustBeLessOrEqual.With(10)},
		},
	}

	for setup, subSetup := range testCases {
		for excl, values := range subSetup {
			testMax := MaxFloat32(setup, excl)

			for input, want := range values {
				t.Run(fmt.Sprintf("%f/%v/%f", setup, excl, input), func(t *testing.T) {
					got := testMax(input)

					if !errors.Is(got, want) {
						t.Error("got", got)
						t.Error("want", want)
					}
				})
			}
		}
	}
}

func TestMaxFloat64(t *testing.T) {
	testCases := map[float64]map[bool]map[float64]error{
		10: {
			true:  {0: nil, 1: nil, 2: nil, 10: ErrMustBeLess.With(10), 100: ErrMustBeLess.With(10)},
			false: {0: nil, 1: nil, 2: nil, 10: nil, 100: ErrMustBeLessOrEqual.With(10)},
		},
	}

	for setup, subSetup := range testCases {
		for excl, values := range subSetup {
			testMax := MaxFloat64(setup, excl)

			for input, want := range values {
				t.Run(fmt.Sprintf("%f/%v/%f", setup, excl, input), func(t *testing.T) {
					got := testMax(input)

					if !errors.Is(got, want) {
						t.Error("got", got)
						t.Error("want", want)
					}
				})
			}
		}
	}
}

func TestMin(t *testing.T) {
	testCases := map[int]map[bool]map[int]error{
		10: {
			true:  {0: ErrMustBeGreater.With(10), 1: ErrMustBeGreater.With(10), 2: ErrMustBeGreater.With(10), 10: ErrMustBeGreater.With(10), 100: nil},
			false: {0: ErrMustBeGreaterOrEqual.With(10), 1: ErrMustBeGreaterOrEqual.With(10), 2: ErrMustBeGreaterOrEqual.With(10), 10: nil, 100: nil},
		},
	}

	for setup, subSetup := range testCases {
		for excl, values := range subSetup {
			testMin := Min(setup, excl)

			for input, want := range values {
				t.Run(fmt.Sprintf("%d/%v/%d", setup, excl, input), func(t *testing.T) {
					got := testMin(input)

					if !errors.Is(got, want) {
						t.Error("got", got)
						t.Error("want", want)
					}
				})
			}
		}
	}
}

func TestMinFloat32(t *testing.T) {
	testCases := map[float32]map[bool]map[float32]error{
		10: {
			true:  {0: ErrMustBeGreater.With(10), 1: ErrMustBeGreater.With(10), 2: ErrMustBeGreater.With(10), 10: ErrMustBeGreater.With(10), 100: nil},
			false: {0: ErrMustBeGreaterOrEqual.With(10), 1: ErrMustBeGreaterOrEqual.With(10), 2: ErrMustBeGreaterOrEqual.With(10), 10: nil, 100: nil},
		},
	}

	for setup, subSetup := range testCases {
		for excl, values := range subSetup {
			testMin := MinFloat32(setup, excl)

			for input, want := range values {
				t.Run(fmt.Sprintf("%f/%v/%f", setup, excl, input), func(t *testing.T) {
					got := testMin(input)

					if !errors.Is(got, want) {
						t.Error("got", got)
						t.Error("want", want)
					}
				})
			}
		}
	}
}

func TestMinFloat64(t *testing.T) {
	testCases := map[float64]map[bool]map[float64]error{
		10: {
			true:  {0: ErrMustBeGreater.With(10), 1: ErrMustBeGreater.With(10), 2: ErrMustBeGreater.With(10), 10: ErrMustBeGreater.With(10), 100: nil},
			false: {0: ErrMustBeGreaterOrEqual.With(10), 1: ErrMustBeGreaterOrEqual.With(10), 2: ErrMustBeGreaterOrEqual.With(10), 10: nil, 100: nil},
		},
	}

	for setup, subSetup := range testCases {
		for excl, values := range subSetup {
			testMin := MinFloat64(setup, excl)

			for input, want := range values {
				t.Run(fmt.Sprintf("%f/%v/%f", setup, excl, input), func(t *testing.T) {
					got := testMin(input)

					if !errors.Is(got, want) {
						t.Error("got", got)
						t.Error("want", want)
					}
				})
			}
		}
	}
}
