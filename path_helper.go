package main

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

// PathHelper represents the application path-helper. Takes a configuration as input, and uses local
// attributes to keep list of files and directories to compose PATH.
type PathHelper struct {
	config      *Config  // parsed command-line flags
	files       []string // slice of files in path.d
	directories []string // directories that will compose PATH
}

// logger for path-helper instance, skip printing when verbose is off.
func (p *PathHelper) logger(format string, v ...interface{}) {
	if p.config.Verbose {
		fmt.Printf(fmt.Sprintf("# %s\n", format), v...)
	}
}

// append a direcotry in global list, making sure it skips duplicates when setting is enabled.
func (p *PathHelper) append(directory string) {
	if p.config.SkipDuplicates {
		for _, d := range p.directories {
			if d == directory {
				p.logger("[WARN] Skipping entry '%s', is already defined.", directory)
				return
			}
		}
	}
	p.directories = append(p.directories, directory)
}

// globPathFiles load list of files in base directory. Returns errors when base directory does not
// exist or when having issues to execute globing.
func (p *PathHelper) globPathFiles() error {
	baseDir := p.config.BaseDir
	p.logger("Inspecting paths directory: '%s'", baseDir)
	if !dirExists(baseDir) {
		return fmt.Errorf("can't find base directory at '%s'", baseDir)
	}

	var err error
	pattern := path.Join(baseDir, "*")
	p.files, err = filepath.Glob(pattern)
	return err
}

// gatherPathDirs based in path files, read and inspect direcotories listed in those. Can return
// errors related to reading files.
func (p *PathHelper) gatherPathDirs() error {
	for _, file := range p.files {
		p.logger("File '%s'", file)
		directories, err := readLines(file)
		if err != nil {
			return fmt.Errorf("can't read file '%s': '%v'", file, err)
		}

		for _, directory := range directories {
			p.logger("\t- '%s'", directory)
			if p.config.SkipNotFound && !dirExists(directory) {
				p.logger("[WARN] Directory '%s' (%s) is not found! Skipping.", directory, file)
				continue
			}
			p.append(directory)
		}
	}
	return nil
}

// pathDirsColonJoined return slice of direcotires joined by colon.
func (p *PathHelper) pathDirsColonJoined() string {
	return strings.Join(p.directories, ":")
}

// RenderExpression print out the shell expression exporting PATH. Will forward errors from methods
// listing and reading path files, and inspecting direcotories present found in those files.
func (p *PathHelper) RenderExpression() (string, error) {
	if err := p.globPathFiles(); err != nil {
		return "", err
	}
	if err := p.gatherPathDirs(); err != nil {
		return "", err
	}

	return fmt.Sprintf("export PATH=\"%s\"", p.pathDirsColonJoined()), nil
}

// NewPathHelper instantiate a PathHelper type.
func NewPathHelper(config *Config) *PathHelper {
	return &PathHelper{
		config:      config,
		files:       []string{},
		directories: []string{},
	}
}
