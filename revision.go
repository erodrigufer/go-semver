// go-sember gets version information embedded during the build process in the
// executable, e.g. revision and revision status from the VCS.
//
// Dependencies: Some of the information being retrieved from the executable,
// can only be accessed from a build, if it was compiled with Go +1.18, most
// likely the compiler will not even allow a build if the version does not
// support the data structures required to access the build information.
//
// This codebase has been modified from its original form, taken from
// autostrada.dev aka. Alex Edwards (2022).
// go-semver Eduardo Rodriguez (@erodrigufer) (c) 2022 -- MIT License.
package gosemver

import (
	"fmt"
	"runtime/debug"
)

// GetRevision, returns a string with the commit hash used to build the
// executable. If the executable was built from files with uncomitted
// modifications related to the last commit, the suffix '-dirty' will be added
// to the revision returned. If no revision information can be retrieved, the
// method returns the string 'unavailable' and an error.
func GetRevision() (string, error) {
	// Last commit's hash used to build the executable.
	var revision string
	// If there  have been uncomitted modifications in the codebase since the
	// last commit, then this variable will be true.
	var modified bool

	bi, ok := debug.ReadBuildInfo()
	if ok {
		for _, s := range bi.Settings {
			switch s.Key {
			case "vcs.revision":
				revision = s.Value
			case "vcs.modified":
				if s.Value == "true" {
					modified = true
				}
			}
		}
	}

	if revision == "" {
		return "unavailable", fmt.Errorf("error: no revision information could be retrieved from the executable.")
	}

	// Add the '-dirty' suffix if the codebase used to build the executable
	// has uncommitted modifications, since the last commit.
	if modified {
		return fmt.Sprintf("%s-dirty", revision), nil
	}

	return revision, nil
}
