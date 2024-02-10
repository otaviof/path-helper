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
func commandLineParser() *pathhelper.Config {
	flag.Usage = func() {
		fmt.Printf(`## path-helper

Helper command-line application to compose "PATH" and "MANPATH" based in a
"paths.d" directory, respecting order of its files and toggles to skip entries.

Environment variables may be used inside the path-files, and these will be
expanded before rendering the final "PATH" and "MANPATH".

To export new "PATH" and "MANPATH" to your shell current instance, run "eval"
against "path-helper" output.

# Usage:
  $ path-helper [-h|--help|flags]

# Examples:
  $ path-helper -v
  $ path-helper -v -s=false -d=false

# Shell-Export:
  $ eval "$(path-helper -s=false -d=false)"
  $ echo ${PATH}
  $ echo ${MANPATH}

# Command-Line Options:
`)
		flag.PrintDefaults()
	}

	cfg := pathhelper.NewConfigFromFlags()
	flag.Parse()
	return cfg
}

func main() {
	cfg := commandLineParser()
	p := pathhelper.NewPathHelper(cfg)
	expr, err := p.RenderExpression()
	if err != nil {
		fatal(err)
	}
	fmt.Println(expr)
}
