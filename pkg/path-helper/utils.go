package pathhelper

import (
	"os"
	"strings"
)

// dirExists check if a directory exists.
func dirExists(dirPath string) bool {
	_, err := os.Stat(dirPath)
	return !os.IsNotExist(err)
}

// readLines read lines from file. Return error in case of having issues to read
// given file.
func readLines(filePath string) ([]string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, nil
	}
	return strings.Split(strings.TrimRight(string(data), "\n"), "\n"), nil
}
