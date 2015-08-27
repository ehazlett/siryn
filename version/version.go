package version

var (
	Version = "0.1.0"

	// GITCOMMIT will be overwritten automatically by the build system
	GitCommit = "HEAD"

	FullVersion = Version + " (" + GitCommit + ")"
)
