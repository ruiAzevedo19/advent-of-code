package solution

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zimmski/osutil"
)

func TestCeresSearch(t *testing.T) {
	type testCase struct {
		Name string

		FilePath string

		ExpectedResult *Result
	}

	testDataPath := filepath.Join("..", "testdata")

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			workspace := t.TempDir()

			testDataFilePath := filepath.Join(testDataPath, tc.FilePath)
			filePath := filepath.Join(workspace, tc.FilePath)
			require.NoError(t, osutil.CopyFile(testDataFilePath, filePath))

			actualResult, actualErr := CeresSearch(filePath)
			require.NoError(t, actualErr)

			assert.Equal(t, tc.ExpectedResult, actualResult)
		})
	}

	validate(t, &testCase{
		Name: "Simple",

		FilePath: "simple.txt",

		ExpectedResult: &Result{
			NumberOfWords:  18,
			NumberOfXWords: 9,
		},
	})
}
