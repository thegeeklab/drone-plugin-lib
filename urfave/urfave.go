// Copyright (c) 2019, Drone Plugins project authors
// Copyright (c) 2021, Robert Kaussow <mail@thegeeklab.de>

// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/thegeeklab/drone-plugin-lib/drone"
	"github.com/urfave/cli/v2"
)

var (
	FlagsBuildCategory      = "Drone Build Flags"
	FlagsRepoCategory       = "Drone Repo Flags"
	FlagsCommitCategory     = "Drone Commit Flags"
	FlagsStageCategory      = "Drone Stage Flags"
	FlagsStepCategory       = "Drone Step Flags"
	FlagsVersioningCategory = "Drone Versioning Flags"
	FlagsSystemCategory     = "Drone System Flags"
	FlagsPluginCategory     = "Plugin Flags"
)

// Flags has the cli.Flags for the Drone plugin.
func Flags() []cli.Flag {
	flags := []cli.Flag{}

	flags = append(flags, buildFlags(FlagsBuildCategory)...)
	flags = append(flags, repoFlags(FlagsRepoCategory)...)
	flags = append(flags, commitFlags(FlagsCommitCategory)...)
	flags = append(flags, stageFlags(FlagsStageCategory)...)
	flags = append(flags, stepFlags(FlagsStepCategory)...)
	flags = append(flags, semVerFlags(FlagsVersioningCategory)...)
	flags = append(flags, calVerFlags(FlagsVersioningCategory)...)
	flags = append(flags, systemFlags(FlagsSystemCategory)...)
	flags = append(flags, networkFlags(FlagsPluginCategory)...)
	flags = append(flags, loggingFlags(FlagsPluginCategory)...)

	return flags
}

// PipelineFromContext creates a drone.Pipeline from the cli.Context.
func PipelineFromContext(ctx *cli.Context) drone.Pipeline {
	return drone.Pipeline{
		Build:  buildFromContext(ctx),
		Repo:   repoFromContext(ctx),
		Commit: commitFromContext(ctx),
		Stage:  stageFromContext(ctx),
		Step:   stepFromContext(ctx),
		SemVer: semVerFromContext(ctx),
		CalVer: calVerFromContext(ctx),
		System: systemFromContext(ctx),
	}
}
