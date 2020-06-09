package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var points = PointList{
	List: []Point{
		{0, 0}, {1, 1}, {1, 1}, {0.1, 0.1}, {-1, 2},
		{1000, -3000}, {5, 0}, {0, -5},
	},
}

func TestPointList_GetSortedWithinDistance(t *testing.T) {
	cases := map[string]struct {
		from     Point
		distance float64
		expected *PointList
	}{
		"far away": {
			from:     Point{500000, 500000},
			distance: 1,
			expected: &PointList{[]Point{}},
		},
		"at a point": {
			from:     Point{0, 0},
			distance: 0,
			expected: &PointList{
				List: []Point{
					{0, 0},
				},
			},
		},
		"duplicate at a point": {
			from:     Point{1.01, 1.01},
			distance: 0.3,
			expected: &PointList{
				List: []Point{
					{1, 1},
					{1, 1},
				},
			},
		},
		"huge distance, all points": {
			from:     Point{0, 0},
			distance: 100000,
			expected: &PointList{
				List: []Point{
					{0, 0}, {0.1, 0.1}, {1, 1}, {1, 1}, {-1, 2}, {5, 0},
					{0, -5}, {1000, -3000},
				},
			},
		},
		"limited distance, order matters": {
			from:     Point{2, 3},
			distance: 5.1,
			expected: &PointList{
				List: []Point{
					{1, 1}, {1, 1}, {-1, 2}, {0.1, 0.1}, {0, 0},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.expected, points.GetSortedWithinDistance(tc.from, tc.distance))
		})
	}
}
