// Copyright (c) 2019, Drone Plugins project authors
// Copyright (c) 2021, Robert Kaussow <mail@thegeeklab.de>

// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

// System represents the available system variables.
type System struct {
	// Proto for the system protocol.
	Proto string

	// Host for the system host name.
	Host string

	// Version for the system version.
	Version string
}

func (s System) String() string {
	return s.Host
}
