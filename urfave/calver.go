// Copyright (c) 2020, the Drone Plugins project authors.

package urfave

import (
	"github.com/thegeeklab/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

// calVerFlags has the cli.Flags for the drone.CalVer.
func calVerFlags(category string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "calver.version",
			Usage: "calver version",
			EnvVars: []string{
				"DRONE_CALVER",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "calver.major",
			Usage: "calver major",
			EnvVars: []string{
				"DRONE_CALVER_MAJOR",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "calver.minor",
			Usage: "calver minor",
			EnvVars: []string{
				"DRONE_CALVER_MINOR",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "calver.micro",
			Usage: "calver micro",
			EnvVars: []string{
				"DRONE_CALVER_MICRO",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "calver.modifier",
			Usage: "calver modifier",
			EnvVars: []string{
				"DRONE_CALVER_MODIFIER",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "calver.short",
			Usage: "calver short",
			EnvVars: []string{
				"DRONE_CALVER_SHORT",
			},
			Category: category,
		},
	}
}

// calVerFromContext creates a drone.CalVer from the cli.Context.
func calVerFromContext(ctx *cli.Context) drone.CalVer {
	return drone.CalVer{
		Version:  ctx.String("calver.version"),
		Major:    ctx.String("calver.major"),
		Minor:    ctx.String("calver.minor"),
		Micro:    ctx.String("calver.micro"),
		Modifier: ctx.String("calver.modifier"),
		Short:    ctx.String("calver.short"),
	}
}
