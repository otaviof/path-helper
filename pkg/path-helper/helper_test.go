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
		ManBaseDir:     "../../test/paths.d",
		PathBaseDir:    "../../test/paths.d",
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
	files, err := p.globPathFiles(c.PathBaseDir)
	t.Logf("Error: '%#v", err)
	assert.NoError(t, err)
	t.Logf("Files: '%#v'", files)
	assert.True(t, len(files) >= expectedLen)

	directories, err := p.inspectPathDirectories(files)
	t.Logf("Error: '%#v", err)
	assert.NoError(t, err)
	t.Logf("Directories: '%#v'", directories)
	assert.True(t, len(directories) >= expectedLen)

	assert.Equal(t, expectedPaths, p.colonJoin(directories))
}

// assertExpression assert primary objective of this app, the shell expression to export PATH. Method
// parameters share test context, configuration and expected export expression.
func assertExpression(t *testing.T, c *Config, expectedExpression string) {
	p := NewPathHelper(c)
	s, err := p.RenderExpression()
	assert.NoError(t, err)
	t.Logf("Expression: '%s'", s)
	assert.NotEmpty(t, s)
	expr := fmt.Sprintf(`PATH="%s" ; MANPATH="%s" ; export PATH MANPATH ;`,
		expectedExpression, expectedExpression)
	assert.Equal(t, expr, s)
}
