# go-semver

The package **go-semver** (Go [Semantic Versioning](https://semver.org/)) provides functions to handle semantic versioning. It integrates nicely with CLI packages like [urfave/cli](https://github.com/urfave/cli) or [cobra](https://github.com/spf13/cobra) to show the current version and/or revision a Go executable.

## Minimum Go version required for a build
In order for the methods in this package to work, the minimum Go version used to compile an executable  must be **Go 1.18** since the struct type _[BuildSetting](https://pkg.go.dev/runtime/debug#BuildSetting)_ was first added in Go 1.18. Otherwise, the commit hash (i. e. revision) and revision status used for the build is not getting embedded into the executable and the executable is not able to access the _BuildSetting_ struct.

If the package [runtime/debug](https://pkg.go.dev/runtime/debug) gets extended in the future, it might even be possible to read the tag with the latest release version used to compile an executable, but as of Go 1.19, this is only achieved when installing a package by running `go install` (one can access the versioning information embedded into an executable by running `go version -m <PATH>` with the path to the executable). 
