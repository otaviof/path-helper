package main

import (
	"flag"
	"fmt"
	"os"

	pathhelper "github.com/otaviof/path-helper/pkg/path-helper"
)

// fatal print out error as a shell comment and exit on error.
func fatal(err error) {
	fmt.Fprintf(os.Stderr, fmt.Sprintf("# [ERROR] %v\n", err))
	os.Exit(1)
}

// commandLineParser handle command line flags, display help message.
func commandLineParser(config *pathhelper.Config) {
	flag.Usage = func() {
		fmt.Printf(`## path-helper

Helper command-line application to compose "PATH" expression based in a "paths.d"
directory, respecting order of files and adding toggles to skip entries.

To export new "PATH" to your shell instance, run "eval" against "path-helper"
output. Examples below.

Usage:
  $ path-helper [-h|--help|flags]

Examples:
  $ path-helper -v
  $ path-helper -v -s=false -d=false

Shell-Export:
  $ eval "$(path-helper -s=false -d=false)"
  $ echo $PATH

Command-Line Options:
`)
		flag.PrintDefaults()
	}

	flag.StringVar(&config.PathBaseDir, "p", "/etc/paths.d", "Paths directory")
	flag.StringVar(&config.ManBaseDir, "m", "/etc/manpaths.d", "Man pages directory")
	flag.BoolVar(&config.SkipNotFound, "d", true, "Skip not found directories")
	flag.BoolVar(&config.SkipDuplicates, "s", true, "Skip duplicated entries")
	flag.BoolVar(&config.Verbose, "v", false, "Verbose")

	flag.Parse()
}

func main() {
	config := &pathhelper.Config{}
	commandLineParser(config)

	p := pathhelper.NewPathHelper(config)
	expr, err := p.RenderExpression()
	if err != nil {
		fatal(err)
	}
	fmt.Println(expr)
}
