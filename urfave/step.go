// Copyright (c) 2019, Drone Plugins project authors
// Copyright (c) 2021, Robert Kaussow <mail@thegeeklab.de>

// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/thegeeklab/drone-plugin-lib/v2/drone"
	"github.com/urfave/cli/v2"
)

// stepFlags has the cli.Flags for the drone.Step.
func stepFlags(category string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "step.name",
			Usage: "step name",
			EnvVars: []string{
				"DRONE_STEP_NAME",
			},
			Category: category,
		},
		&cli.IntFlag{
			Name:  "step.number",
			Usage: "step number",
			EnvVars: []string{
				"DRONE_STEP_NUMBER",
			},
			Category: category,
		},
	}
}

// stepFromContext creates a drone.Step from the cli.Context.
func stepFromContext(ctx *cli.Context) drone.Step {
	return drone.Step{
		Name:   ctx.String("step.name"),
		Number: ctx.Int("step.number"),
	}
}
