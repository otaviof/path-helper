package pathhelper

// Config all configuration entries supported by this application.
type Config struct {
	PathBaseDir    string // paths.d directory path
	ManBaseDir     string // manpaths.d directory path
	SkipNotFound   bool   // make sure all directories in path exist
	SkipDuplicates bool   // skip duplicated entries to compose PATH
	Verbose        bool   // make output verbose
}
