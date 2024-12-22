package solution

import (
	"path/filepath"
	"testing"

	challengetesting "advent-of-code/challenge/testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDiskFragmenter(t *testing.T) {
	validate := func(t *testing.T, tc *challengetesting.TestCaseChallenge) {
		tc.Validate(t)
	}

	validate(t, &challengetesting.TestCaseChallenge{
		Name: "Simple",

		FilePath: filepath.Join("testdata", "simple.txt"),
		Solution: DiskFragmenter,

		ExpectedResult: &Result{
			Checksum:          1928,
			ChecksumCompacted: 2858,
		},
	})
}

func TestLooseFormat(t *testing.T) {
	type testCase struct {
		Name string

		DiskMap string

		ExpectedLooseFormat []string
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualLooseFormat, actualErr := looseFormat(tc.DiskMap)
			require.NoError(t, actualErr)

			assert.Equal(t, tc.ExpectedLooseFormat, actualLooseFormat)
		})
	}

	validate(t, &testCase{
		Name: "Simple",

		DiskMap: "12345",

		ExpectedLooseFormat: []string{
			"0", ".", ".", "1", "1", "1", ".", ".", ".", ".", "2", "2", "2", "2", "2",
		},
	})
	validate(t, &testCase{
		Name: "Complex",

		DiskMap: "2333133121414131402",

		ExpectedLooseFormat: []string{
			"0", "0", ".", ".", ".", "1", "1", "1", ".", ".", ".", "2", ".", ".", ".", "3", "3", "3", ".", "4", "4", ".", "5", "5", "5", "5", ".", "6", "6", "6", "6", ".", "7", "7", "7", ".", "8", "8", "8", "8", "9", "9",
		},
	})
}

func TestMoveFileBlocks(t *testing.T) {
	type testCase struct {
		Name string

		FileBlocks []string

		ExpectedFileBlocks []string
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualFileBlocks := moveFileBlocks(tc.FileBlocks)

			assert.Equal(t, tc.ExpectedFileBlocks, actualFileBlocks)
		})
	}

	validate(t, &testCase{
		Name: "Simple",

		FileBlocks: []string{
			"0", ".", ".", "1", "1", "1", ".", ".", ".", ".", ".", "2", "2", "2", "2", "2",
		},
		ExpectedFileBlocks: []string{
			"0", "2", "2", "1", "1", "1", "2", "2", "2", ".", ".", ".", ".", ".", ".", ".",
		},
	})
	validate(t, &testCase{
		Name: "Complex",

		FileBlocks: []string{
			"0", "0", ".", ".", ".", "1", "1", "1", ".", ".", ".", "2", ".", ".", ".", "3", "3", "3", ".", "4", "4", ".", "5", "5", "5", "5", ".", "6", "6", "6", "6", ".", "7", "7", "7", ".", "8", "8", "8", "8", "9", "9",
		},

		ExpectedFileBlocks: []string{
			"0", "0", "9", "9", "8", "1", "1", "1", "8", "8", "8", "2", "7", "7", "7", "3", "3", "3", "6", "4", "4", "6", "5", "5", "5", "5", "6", "6", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".", ".",
		},
	})
}

func TestCompactFiles(t *testing.T) {
	type testCase struct {
		Name string

		FileBlocks []string

		ExpectedFileBlocks []string
	}

	validate := func(t *testing.T, tc *testCase) {
		t.Run(tc.Name, func(t *testing.T) {
			actualFileBlocks := compactFiles(tc.FileBlocks)

			assert.Equal(t, tc.ExpectedFileBlocks, actualFileBlocks)
		})
	}

	validate(t, &testCase{
		Name: "Complex",

		FileBlocks: []string{
			"0", "0", ".", ".", ".", "1", "1", "1", ".", ".", ".", "2", ".", ".", ".", "3", "3", "3", ".", "4", "4", ".", "5", "5", "5", "5", ".", "6", "6", "6", "6", ".", "7", "7", "7", ".", "8", "8", "8", "8", "9", "9",
		},

		ExpectedFileBlocks: []string{
			"0", "0", "9", "9", "2", "1", "1", "1", "7", "7", "7", ".", "4", "4", ".", "3", "3", "3", ".", ".", ".", ".", "5", "5", "5", "5", ".", "6", "6", "6", "6", ".", ".", ".", ".", ".", "8", "8", "8", "8", ".", ".",
		},
	})
}
