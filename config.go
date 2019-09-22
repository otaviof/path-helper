package main

// Config all configuration entries supported by this application.
type Config struct {
	BaseDir        string // base directory
	SkipNotFound   bool   // make sure all directories in path exist
	SkipDuplicates bool   // skip duplicated entries to compose PATH
	Verbose        bool   // make output verbose
}
