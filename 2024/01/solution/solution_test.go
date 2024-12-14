package solution

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zimmski/osutil"
)

func TestHistoriansHysteria(t *testing.T) {
	type testCase struct {
		Name string

		FilePath string

		ExpectedResult *Result
	}

	testDataFilePath := filepath.Join("..", "testdata")

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			workspace := t.TempDir()
			filePath := filepath.Join(workspace, tc.FilePath)

			require.NoError(t, osutil.CopyFile(filepath.Join(testDataFilePath, tc.FilePath), filePath))

			actualResult, actualErr := HistorianHysteria(filePath)
			require.NoError(t, actualErr)

			assert.Equal(t, tc.ExpectedResult, actualResult)
		})
	}

	validate(t, &testCase{
		Name: "Simple",

		FilePath: "simple.txt",

		ExpectedResult: &Result{
			TotalDistance:   11,
			SimilarityScore: 31,
		},
	})
}
