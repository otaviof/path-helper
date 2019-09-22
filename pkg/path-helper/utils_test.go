package pathhelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtilsDirExists(t *testing.T) {
	assert.True(t, dirExists("."))
	assert.False(t, dirExists("/x/y/z/path-helper"))
}

func TestUtilsReadLines(t *testing.T) {
	lines, err := readLines("utils_test.go")
	assert.NoError(t, err)
	assert.True(t, len(lines) > 0)
}
