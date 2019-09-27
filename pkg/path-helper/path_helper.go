package pathhelper

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

// PathHelper represents the application path-helper. Takes a configuration as input, and uses local
// attributes to keep list of files and directories to compose PATH.
type PathHelper struct {
	config *Config // parsed command-line flags
}

// logger for path-helper instance, skip printing when verbose is off.
func (p *PathHelper) logger(format string, v ...interface{}) {
	if p.config.Verbose {
		fmt.Printf(fmt.Sprintf("# %s\n", format), v...)
	}
}

// append a directory in global list, making sure it skips duplicates when setting is enabled.
func (p *PathHelper) append(directories []string, directory string) []string {
	if p.config.SkipDuplicates {
		for _, d := range directories {
			if d == directory {
				p.logger("[WARN] Skipping entry '%s', is already defined.", directory)
				return directories
			}
		}
	}
	return append(directories, directory)
}

// globPathFiles load list of files in base directory. Returns errors when base directory does not
// exist or when having issues to execute globing.
func (p *PathHelper) globPathFiles(baseDir string) ([]string, error) {
	p.logger("Inspecting paths directory: '%s'", baseDir)
	if !dirExists(baseDir) {
		return nil, fmt.Errorf("can't find base directory at '%s'", baseDir)
	}

	pattern := path.Join(baseDir, "*")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	return files, nil
}

// inspectPathDirectories based in path files, read and inspect direcotories listed in those. Can
// return errors related to reading files.
func (p *PathHelper) inspectPathDirectories(files []string) ([]string, error) {
	directories := []string{}
	for _, file := range files {
		p.logger("File '%s'", file)
		lines, err := readLines(file)
		if err != nil {
			return nil, fmt.Errorf("can't read file '%s': '%v'", file, err)
		}

		for _, directory := range lines {
			p.logger("\t- '%s'", directory)
			if strings.HasPrefix(directory, "#") {
				continue
			}
			if p.config.SkipNotFound && !dirExists(directory) {
				p.logger("[WARN] Directory '%s' (%s) is not found! Skipping.", directory, file)
				continue
			}
			directories = p.append(directories, directory)
		}
	}
	return directories, nil
}

// collect glob for files and open them to extract contents. File contents are threated as path
// directories, therefore configuration directive applies on them. It can return error when having
// problems to glob directories and on reading files.
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

// pathDirsColonJoined return slice of direcotires joined by colon.
func (p *PathHelper) colonJoin(directories []string) string {
	return strings.Join(directories, ":")
}

// RenderExpression print out the shell expression exporting PATH. Will forward errors from methods
// listing and reading path files, and inspecting direcotories present found in those files.
func (p *PathHelper) RenderExpression() (string, error) {
	pathDirectories, err := p.collect(p.config.PathBaseDir)
	if err != nil {
		return "", err
	}

	manDirectories, err := p.collect(p.config.ManBaseDir)
	if err != nil {
		return "", err
	}

	expr := fmt.Sprintf(`PATH="%s" ; MANPATH="%s" ; export PATH MANPATH ;`,
		p.colonJoin(pathDirectories), p.colonJoin(manDirectories))
	return expr, nil
}

// NewPathHelper instantiate a PathHelper type.
func NewPathHelper(config *Config) *PathHelper {
	return &PathHelper{config: config}
}
