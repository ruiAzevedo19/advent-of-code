package solution

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zimmski/osutil"
)

func TestParse(t *testing.T) {
	type testCase struct {
		Name string

		FilePath string

		ExpectedGrid []string
	}

	testDataPath := "testdata"

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			workspace := t.TempDir()

			testDataFilePath := filepath.Join(testDataPath, tc.FilePath)
			filePath := filepath.Join(workspace, tc.FilePath)
			require.NoError(t, osutil.CopyFile(testDataFilePath, filePath))

			actualGrid, actualErr := Parse(filePath)
			require.NoError(t, actualErr)

			assert.Equal(t, tc.ExpectedGrid, actualGrid)
		})
	}

	validate(t, &testCase{
		Name: "Simple",

		FilePath: "simple.txt",

		ExpectedGrid: []string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX",
		},
	})
}
