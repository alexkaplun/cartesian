package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint_Distance(t *testing.T) {
	cases := map[string]struct {
		from     Point
		to       Point
		expected float64
	}{
		"zero": {
			from:     Point{0, 0},
			to:       Point{0, 0},
			expected: 0,
		},
		"same point": {
			from:     Point{-5.2, 4.06},
			to:       Point{-5.2, 4.06},
			expected: 0,
		},
		"1": {
			from:     Point{0, 0},
			to:       Point{4.3, 5.8},
			expected: 10.1,
		},
		"2": {
			from:     Point{-2, 1},
			to:       Point{2.5, -3.08},
			expected: 8.58,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, tc.from.Distance(tc.to))
		})
	}
}
