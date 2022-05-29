// Copyright (c) 2019, Drone Plugins project authors
// Copyright (c) 2021, Robert Kaussow <mail@thegeeklab.de>

// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

// loggingFlags has the cli.Flags for logging config.
func loggingFlags(category string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "log-level",
			Usage:    "log level",
			EnvVars:  []string{"PLUGIN_LOG_LEVEL"},
			Category: category,
		},
	}
}

// LoggingFromContext sets the logrus logging level.
func LoggingFromContext(ctx *cli.Context) {
	lvl, err := logrus.ParseLevel(ctx.String("log-level"))
	if err != nil {
		lvl = logrus.InfoLevel
	}

	logrus.SetLevel(lvl)
}
