// Copyright (c) 2019, Drone Plugins project authors
// Copyright (c) 2021, Robert Kaussow <mail@thegeeklab.de>

// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/thegeeklab/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

// systemFlags has the cli.Flags for the drone.System.
func systemFlags(category string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "system.proto",
			Usage: "system proto",
			EnvVars: []string{
				"DRONE_SYSTEM_PROTO",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "system.host",
			Usage: "system host",
			EnvVars: []string{
				"DRONE_SYSTEM_HOST",
				"DRONE_SYSTEM_HOSTNAME",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "system.version",
			Usage: "system version",
			EnvVars: []string{
				"DRONE_SYSTEM_VERSION",
			},
			Category: category,
		},
	}
}

// systemFromContext creates a drone.System from the cli.Context.
func systemFromContext(ctx *cli.Context) drone.System {
	return drone.System{
		Proto:   ctx.String("system.proto"),
		Host:    ctx.String("system.host"),
		Version: ctx.String("system.version"),
	}
}
