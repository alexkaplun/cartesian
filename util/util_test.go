package util

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/alexkaplun/cartesian/model"
)

const TEST_FILES_DIR = "../data/"

func TestLoadPointListFromCsv(t *testing.T) {
	cases := map[string]struct {
		fileName string
		isError  bool
		expected *model.PointList
	}{
		"missing file": {
			fileName: "i am a missing file",
			isError:  true,
		},
		"wrong json format": {
			fileName: "test1.json",
			isError:  true,
		},
		"empty points": {
			fileName: "test2.json",
			isError:  false,
			expected: &model.PointList{
				List: []model.Point{},
			},
		},
		"valid data": {
			fileName: "test3.json",
			isError:  false,
			expected: &model.PointList{
				List: []model.Point{
					{X: 1, Y: 1},
					{X: 7, Y: 0},
					{X: 0, Y: 0},
					{X: 0, Y: 0.1},
					{X: -0.1, Y: -17.988},
				},
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			if tc.isError {
				_, err := LoadPointListFromCsv(TEST_FILES_DIR + tc.fileName)
				require.Error(t, err)
			} else {
				list, err := LoadPointListFromCsv(TEST_FILES_DIR + tc.fileName)
				require.NoError(t, err)
				assert.EqualValues(t, tc.expected, list)
			}
		})
	}
}
