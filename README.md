# go-semver

The package **go-semver** (Go [Semantic Versioning](https://semver.org/)) provides functions to handle semantic versioning. It integrates nicely with CLI packages like [urfave/cli](https://github.com/urfave/cli) or [cobra](https://github.com/spf13/cobra) to show the current version and/or revision of a Go executable.

## How to use
1. Add this package as a dependency to your Go project by running the command `go get github.com/erodrigufer/go-semver` inside the directory where the go.mod file of your project lies.
2. Import the package into the Go files that will use its methods.

### Show the build revision of a Go executable
The method `GetRevision()` returns the commit's hash (revision) of the last commit of the codebase used to build an executable. If there are uncommitted modifications in the codebase since the last commit, the method will add the suffix '-dirty' to the string returned with the commit's hash.

If the method was unable to retrieve a revision, it will return the string 'unavailable' and a non-nil error.

## Troubleshooting
### Revision is unavailable, i.e. is not showing up
You have to build your executable from within the VCS (Version Control System, e.g. git) repository/codebase in order for the `go` build tool to get the build information related to the VCS, like the revision, or its status (was it built with uncommitted modifications related to the last commit).

**If you build your executable outside a VCS repository (aka not within a git repo) then the build tool is not going to be able to retrieve any revision or version information, and the `GetRevision()` function will return an error.**

### Minimum Go version required for a build
In order for the methods in this package to work, the minimum Go version used to compile an executable  must be **Go 1.18** since the struct type _[BuildSetting](https://pkg.go.dev/runtime/debug#BuildSetting)_ was first added in Go 1.18. Otherwise, the commit hash (i. e. revision) and revision status used for the build is not getting embedded into the executable and the executable is not able to access the _BuildSetting_ struct.

If the package [runtime/debug](https://pkg.go.dev/runtime/debug) gets extended in the future, it might even be possible to read the tag with the latest release version used to compile an executable, but as of Go 1.19, this is only achieved when installing a package by running `go install` (one can access the versioning information embedded into an executable by running `go version -m <PATH>` with the path to the executable). 
