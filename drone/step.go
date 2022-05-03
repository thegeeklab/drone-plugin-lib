// Copyright (c) 2019, Drone Plugins project authors
// Copyright (c) 2021, Robert Kaussow <mail@thegeeklab.de>

// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package drone

// Step represents the currently running step within the stage.
type Step struct {
	// Name for the name of the current step.
	Name string

	// Number is the numeric value of the step.
	Number int
}

func (s Step) String() string {
	return s.Name
}
