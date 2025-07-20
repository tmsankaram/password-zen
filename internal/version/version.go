package version

import (
	"fmt"
	"runtime"
)

var (
	// Version is the current version of password-zen
	Version = "1.0.0"
	// BuildDate is the date when the binary was built
	BuildDate = "development"
	// GitCommit is the git commit hash
	GitCommit = "development"
)

// Info returns version information
func Info() string {
	return fmt.Sprintf("Password Zen v%s\nBuilt: %s\nCommit: %s\nGo: %s %s/%s",
		Version,
		BuildDate,
		GitCommit,
		runtime.Version(),
		runtime.GOOS,
		runtime.GOARCH,
	)
}

// Short returns just the version number
func Short() string {
	return Version
}
