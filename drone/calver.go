// Copyright (c) 2020, the Drone Plugins project authors.

package drone

// CalVer represents the calendar version of the currently running build.
//
// This value is only applicable for tags. If the tag cannot be parsed into
// a calendar version then the value will be empty.
type CalVer struct {
	// Version is the full calendar version.
	Version string

	// Major is the major version.
	Major string

	// Minor is the minor version.
	Minor string

	// Micro is the micro version.
	Micro string

	// Modifier is a modifier for the version.
	Modifier string

	// Short is the short version.
	//
	// This does not include the modifier.
	Short string
}

func (c CalVer) String() string {
	return c.Version
}
