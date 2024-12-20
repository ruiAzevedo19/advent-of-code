package parse

import (
	"advent-of-code/2024/05/model"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zimmski/osutil/bytesutil"
)

func TestParse(t *testing.T) {
	type testCase struct {
		Name string

		FileContent string

		ExpectedRules   []*model.Rule
		ExpectedUpdates [][]int
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			workspace := t.TempDir()

			filePath := filepath.Join(workspace, "plain.txt")
			require.NoError(t, os.WriteFile(filePath, []byte(bytesutil.StringTrimIndentations(tc.FileContent)), 0700))

			actualRules, actualUpdates, actualErr := Parse(filePath)
			require.NoError(t, actualErr)

			assert.Equal(t, tc.ExpectedRules, actualRules)
			assert.Equal(t, tc.ExpectedUpdates, actualUpdates)
		})
	}

	validate(t, &testCase{
		Name: "Single",

		FileContent: `
            1|2

            1,2
        `,

		ExpectedRules: []*model.Rule{
			{
				X: 1,
				Y: 2,
			},
		},
		ExpectedUpdates: [][]int{
			{1, 2},
		},
	})
	validate(t, &testCase{
		Name: "Multiple",

		FileContent: `
            1|2
			3|4
			5|6
			7|8
			9|0

            1,2,3,4,5
			6,7,8,9,10
			11,12,13,14,15
			16,17,18,19,20
        `,

		ExpectedRules: []*model.Rule{
			{
				X: 1,
				Y: 2,
			},
			{
				X: 3,
				Y: 4,
			},
			{
				X: 5,
				Y: 6,
			},
			{
				X: 7,
				Y: 8,
			},
			{
				X: 9,
				Y: 0,
			},
		},
		ExpectedUpdates: [][]int{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
		},
	})
}
