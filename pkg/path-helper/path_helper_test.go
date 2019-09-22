package pathhelper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPathHelper run through common use-cases setting different options and asserting.
func TestPathHelper(t *testing.T) {
	c := &Config{
		SkipDuplicates: true,
		SkipNotFound:   true,
		Verbose:        true,
		BaseDir:        "../../test/paths.d",
	}
	expectedJoinedPaths := "/a/a/a:/b/b/b:/c/c/c:/d/d/d"
	expectedJoinedPathsDuplicated := "/a/a/a:/b/b/b:/c/c/c:/d/d/d:/d/d/d"

	t.Run("without-duplicates-and-without-not-founds", func(t *testing.T) {
		t.Run("inspecting-directories-and-files", func(t *testing.T) {
			assertFilesAndDirectories(t, c, 0, "")
		})

		t.Run("RenderExpression", func(t *testing.T) {
			assertExpression(t, c, "")
		})
	})

	t.Run("with-duplicates", func(t *testing.T) {
		c.SkipDuplicates = false
		c.SkipNotFound = true

		t.Run("inspecting-directories-and-files", func(t *testing.T) {
			assertFilesAndDirectories(t, c, 0, "")
		})

		t.Run("RenderExpression", func(t *testing.T) {
			assertExpression(t, c, "")
		})
	})

	t.Run("with-not-founds", func(t *testing.T) {
		c.SkipDuplicates = true
		c.SkipNotFound = false

		t.Run("inspecting-directories-and-files", func(t *testing.T) {
			assertFilesAndDirectories(t, c, 1, expectedJoinedPaths)
		})

		t.Run("RenderExpression", func(t *testing.T) {
			assertExpression(t, c, expectedJoinedPaths)
		})
	})

	t.Run("with-duplicates-and-with-not-founds", func(t *testing.T) {
		c.SkipDuplicates = false
		c.SkipNotFound = false

		t.Run("inspecting-directories-and-files", func(t *testing.T) {
			assertFilesAndDirectories(t, c, 1, expectedJoinedPathsDuplicated)
		})

		t.Run("RenderExpression", func(t *testing.T) {
			assertExpression(t, c, expectedJoinedPathsDuplicated)
		})
	})
}

// assertFilesAndDirectories run assertion of "files" and "directories" attributes. Method parameters
// share the test context and expected characteristics.
func assertFilesAndDirectories(t *testing.T, c *Config, expectedLen int, expectedPaths string) {
	p := NewPathHelper(c)
	err := p.globPathFiles()
	t.Logf("Error: '%#v", err)
	assert.NoError(t, err)
	t.Logf("Files: '%#v'", p.files)
	assert.True(t, len(p.files) >= expectedLen)

	err = p.gatherPathDirs()
	t.Logf("Error: '%#v", err)
	assert.NoError(t, err)
	t.Logf("Directories: '%#v'", p.directories)
	assert.True(t, len(p.directories) >= expectedLen)

	assert.Equal(t, expectedPaths, p.pathDirsColonJoined())
}

// assertExpression assert primary objective of this app, the shell expression to export PATH. Method
// parameters share test context, configuration and expected export expression.
func assertExpression(t *testing.T, c *Config, expectedExpression string) {
	p := NewPathHelper(c)
	s, err := p.RenderExpression()
	assert.NoError(t, err)
	t.Logf("Expression: '%s'", s)
	assert.NotEmpty(t, s)
	assert.Equal(t, fmt.Sprintf("export PATH=\"%s\"", expectedExpression), s)
}
