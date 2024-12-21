package testing

import (
	"advent-of-code/challenge"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zimmski/osutil"
)

type TestCaseChallenge struct {
	Name string

	FilePath string
	Solution challenge.Solution

	ExpectedResult challenge.Result
}

func (tc *TestCaseChallenge) Validate(t *testing.T) {
	t.Run(tc.Name, func(t *testing.T) {
		workspace := t.TempDir()

		filePath := filepath.Join(workspace, "input.txt")
		require.NoError(t, osutil.CopyFile(tc.FilePath, filePath))

		actualResult, actualErr := tc.Solution(filePath)
		require.NoError(t, actualErr)

		assert.Equal(t, tc.ExpectedResult, actualResult)
	})
}
