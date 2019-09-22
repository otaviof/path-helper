package pathhelper

import (
	"bufio"
	"io"
	"os"
)

// dirExists check if a directory exists.
func dirExists(dirPath string) bool {
	_, err := os.Stat(dirPath)
	return !os.IsNotExist(err)
}

// readLines read lines from file. Return error in case of having issues to read given file.
func readLines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return []string{}, err
	}
	defer f.Close()

	var lines []string
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return []string{}, err
		}
		lines = append(lines, string(line))
	}
	return lines, nil
}
