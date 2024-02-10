package pathhelper

import "flag"

// Config all configuration entries supported by this application.
type Config struct {
	PathBaseDir    string // paths.d directory path
	ManBaseDir     string // manpaths.d directory path
	SkipNotFound   bool   // make sure all directories in path exist
	SkipDuplicates bool   // skip duplicated entries to compose PATH
	Verbose        bool   // make output verbose
}

func NewConfig() *Config {
	return &Config{
		PathBaseDir:    "/etc/paths.d",
		ManBaseDir:     "/etc/manpaths.d",
		SkipNotFound:   true,
		SkipDuplicates: true,
		Verbose:        false,
	}
}

func NewConfigFromFlags() *Config {
	cfg := NewConfig()
	flag.StringVar(&cfg.PathBaseDir, "p", cfg.PathBaseDir, "Paths directory")
	flag.StringVar(&cfg.ManBaseDir, "m", cfg.ManBaseDir, "Man pages directory")
	flag.BoolVar(&cfg.SkipNotFound, "d", cfg.SkipNotFound, "Skip not found directories")
	flag.BoolVar(&cfg.SkipDuplicates, "s", cfg.SkipDuplicates, "Skip duplicated entries")
	flag.BoolVar(&cfg.Verbose, "v", cfg.Verbose, "Verbose")
	return cfg
}
