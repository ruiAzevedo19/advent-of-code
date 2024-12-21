package solution

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zimmski/osutil"
)

func TestPrintQueue(t *testing.T) {
	type testCase struct {
		Name string

		FilePath string

		ExpectedValidUpdates       [][]int
		ExpectedInvalidUpdates     [][]int
		ExpectedMiddlePageSum      int
		ExpectedFixedUpdates       [][]int
		ExpectedFixedMiddlePageSum int
	}

	testDataPath := "testdata"

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			workspace := t.TempDir()

			filePath := filepath.Join(workspace, tc.FilePath)
			require.NoError(t, osutil.CopyFile(filepath.Join(testDataPath, tc.FilePath), filePath))

			result, actualErr := PrintQueue(filePath)
			require.NoError(t, actualErr)
			actualResult, ok := result.(*Result)
			require.True(t, ok)

			assert.Equal(t, tc.ExpectedValidUpdates, actualResult.ValidUpdates)
			assert.Equal(t, tc.ExpectedInvalidUpdates, actualResult.InvalidUpdates)
			assert.Equal(t, tc.ExpectedMiddlePageSum, actualResult.MiddlePageSumValidUpdates)

			assert.Equal(t, tc.ExpectedFixedUpdates, actualResult.FixedUpdates)
			assert.Equal(t, tc.ExpectedFixedMiddlePageSum, actualResult.MiddlePageSumFixedUpdates)
		})
	}

	validate(t, &testCase{
		Name: "Simple",

		FilePath: "simple.txt",

		ExpectedValidUpdates: [][]int{
			{75, 47, 61, 53, 29},
			{97, 61, 53, 29, 13},
			{75, 29, 13},
		},
		ExpectedInvalidUpdates: [][]int{
			{75, 97, 47, 61, 53},
			{61, 13, 29},
			{97, 13, 75, 29, 47},
		},
		ExpectedMiddlePageSum: 143,

		ExpectedFixedUpdates: [][]int{
			{97, 75, 47, 61, 53},
			{61, 29, 13},
			{97, 75, 47, 29, 13},
		},
		ExpectedFixedMiddlePageSum: 123,
	})
}
