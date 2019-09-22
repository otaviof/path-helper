package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathHelper(t *testing.T) {
	config := &Config{Verbose: true, BaseDir: "./test/paths.d"}
	expectedJoinedPaths := "/a/a/a:/b/b/b:/c/c/c:/d/d/d"

	t.Run("without-duplicates", func(t *testing.T) {
		t.Run("inspecting-directories-and-files", func(t *testing.T) {
			inspectPathFilesAndDirectories(t, config, 1, expectedJoinedPaths)
		})

		t.Run("RenderExpression", func(t *testing.T) {
			inspectRenderedExpression(t, config, expectedJoinedPaths)
		})
	})

	t.Run("with-duplicates", func(t *testing.T) {
		config.SkipDuplicates = true

		t.Run("inspecting-directories-and-files", func(t *testing.T) {
			inspectPathFilesAndDirectories(t, config, 1, expectedJoinedPaths)
		})

		t.Run("RenderExpression", func(t *testing.T) {
			inspectRenderedExpression(t, config, expectedJoinedPaths)
		})
	})

	t.Run("with-skip-not-found", func(t *testing.T) {
		config.SkipDuplicates = true
		config.SkipNotFound = true

		t.Run("inspecting-directories-and-files", func(t *testing.T) {
			inspectPathFilesAndDirectories(t, config, 0, "")
		})

		t.Run("RenderExpression", func(t *testing.T) {
			inspectRenderedExpression(t, config, "")
		})
	})
}

func inspectPathFilesAndDirectories(
	t *testing.T,
	config *Config,
	expectedLen int,
	expectedJoinedPaths string,
) {
	p := NewPathHelper(config)
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

	assert.Equal(t, expectedJoinedPaths, p.pathDirsColonJoined())
}

func inspectRenderedExpression(t *testing.T, config *Config, expectedJoinedPaths string) {
	p := NewPathHelper(config)
	s, err := p.RenderExpression()
	assert.NoError(t, err)
	t.Logf("Expression: '%s'", s)
	assert.NotEmpty(t, s)
	assert.Equal(t, fmt.Sprintf("export PATH=\"%s\"", expectedJoinedPaths), s)
}
