package pathhelper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// PathHelper represents the application path-helper. Takes a configuration as
// input, and uses local attributes to keep list of files and directories to
// compose PATH.
type PathHelper struct {
	config *Config             // parsed command-line flags
	seen   map[string]struct{} // tracks duplicates for O(1) lookups
}

// logger for path-helper instance, skip printing when verbose is off.
func (p *PathHelper) logger(format string, v ...any) {
	if p.config.Verbose {
		fmt.Fprintf(os.Stderr, "# "+format+"\n", v...)
	}
}

// append a directory in global list, making sure it skips duplicates when setting
// is enabled.
func (p *PathHelper) append(directories []string, directory string) []string {
	if p.config.SkipDuplicates {
		if _, exists := p.seen[directory]; exists {
			p.logger("[WARN] Skipping entry '%s', is already defined.", directory)
			return directories
		}
		p.seen[directory] = struct{}{}
	}
	return append(directories, directory)
}

// globPathFiles load list of files in base directory. Returns errors when base
// directory does not exist or when having issues to read the directory.
func (p *PathHelper) globPathFiles(baseDir string) ([]string, error) {
	p.logger("Inspecting paths directory: '%s'", baseDir)

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return nil, fmt.Errorf("can't find base directory at '%s': %w", baseDir, err)
	}
	files := make([]string, 0, len(entries))
	for _, e := range entries {
		if !e.IsDir() {
			files = append(files, filepath.Join(baseDir, e.Name()))
		}
	}
	return files, nil
}

// inspectPathDirectories based in path files, read and inspect directories
// listed, and it also expands environment variables. Can return errors related to
// reading files.
func (p *PathHelper) inspectPathDirectories(files []string) ([]string, error) {
	clear(p.seen)
	directories := make([]string, 0, len(files)*4)
	for _, file := range files {
		p.logger("File '%s'", file)
		lines, err := readLines(file)
		if err != nil {
			return nil, fmt.Errorf("can't read file '%s': %w", file, err)
		}

		for _, directory := range lines {
			if strings.HasPrefix(directory, "#") {
				continue
			}
			directory = os.ExpandEnv(directory)
			p.logger("\t- '%s'", directory)
			if p.config.SkipNotFound && !dirExists(directory) {
				p.logger("[WARN] Directory '%s' (%s) is not found! Skipping.",
					directory, file)
				continue
			}
			directories = p.append(directories, directory)
		}
	}
	return directories, nil
}

// collect glob for files and open them to extract contents. File contents are
// treated as path directories, therefore configuration directive applies on them.
// It can return error when having problems to glob directories and on reading
// files.
func (p *PathHelper) collect(baseDir string) ([]string, error) {
	files, err := p.globPathFiles(baseDir)
	if err != nil {
		return nil, err
	}
	directories, err := p.inspectPathDirectories(files)
	if err != nil {
		return nil, err
	}
	return directories, nil
}

// colonJoin return slice of directories joined by colon.
func (p *PathHelper) colonJoin(directories []string) string {
	return strings.Join(directories, ":")
}

// RenderExpression print out the shell expression exporting PATH. Will forward
// errors from methods listing and reading path files, and inspecting directories
// present found in those files.
func (p *PathHelper) RenderExpression() (string, error) {
	pathDirectories, err := p.collect(p.config.PathBaseDir)
	if err != nil {
		return "", err
	}

	manDirectories, err := p.collect(p.config.ManBaseDir)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	b.Grow(128)
	b.WriteString(`PATH="`)
	b.WriteString(p.colonJoin(pathDirectories))
	b.WriteString(`" ; MANPATH="`)
	b.WriteString(p.colonJoin(manDirectories))
	b.WriteString(`" ; export PATH MANPATH ;`)
	return b.String(), nil
}

// NewPathHelper instantiate a PathHelper type.
func NewPathHelper(config *Config) *PathHelper {
	return &PathHelper{
		config: config,
		seen:   make(map[string]struct{}, 32),
	}
}
