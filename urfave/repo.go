// Copyright (c) 2019, Drone Plugins project authors
// Copyright (c) 2021, Robert Kaussow <mail@thegeeklab.de>

// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"github.com/thegeeklab/drone-plugin-lib/v2/drone"
	"github.com/urfave/cli/v2"
)

// repoFlags has the cli.Flags for the drone.Repo
func repoFlags(category string) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  "repo.slug",
			Usage: "repo slug",
			EnvVars: []string{
				"DRONE_REPO",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "repo.scm",
			Usage: "repo scm",
			EnvVars: []string{
				"DRONE_REPO_SCM",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "repo.owner",
			Usage: "repo owner",
			EnvVars: []string{
				"DRONE_REPO_OWNER",
				"DRONE_REPO_NAMESPACE",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "repo.name",
			Usage: "repo name",
			EnvVars: []string{
				"DRONE_REPO_NAME",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "repo.link",
			Usage: "repo link",
			EnvVars: []string{
				"DRONE_REPO_LINK",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "repo.branch",
			Usage: "repo branch",
			EnvVars: []string{
				"DRONE_REPO_BRANCH",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "repo.http-url",
			Usage: "repo http url",
			EnvVars: []string{
				"DRONE_REMOTE_URL",
				"DRONE_GIT_HTTP_URL",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "repo.ssh-url",
			Usage: "repo ssh url",
			EnvVars: []string{
				"DRONE_GIT_SSH_URL",
			},
			Category: category,
		},
		&cli.StringFlag{
			Name:  "repo.visibility",
			Usage: "repo visibility",
			EnvVars: []string{
				"DRONE_REPO_VISIBILITY",
			},
			Category: category,
		},
		&cli.BoolFlag{
			Name:  "repo.private",
			Usage: "repo private",
			EnvVars: []string{
				"DRONE_REPO_PRIVATE",
			},
			Category: category,
		},
	}
}

// repoFromContext creates a drone.Repo from the cli.Context.
func repoFromContext(ctx *cli.Context) drone.Repo {
	return drone.Repo{
		Slug:       ctx.String("repo.slug"),
		SCM:        ctx.String("repo.scm"),
		Owner:      ctx.String("repo.owner"),
		Name:       ctx.String("repo.name"),
		Link:       ctx.String("repo.link"),
		Branch:     ctx.String("repo.branch"),
		HTTPURL:    ctx.String("repo.http-url"),
		SSHURL:     ctx.String("repo.ssh-url"),
		Visibility: ctx.String("repo.visibility"),
		Private:    ctx.Bool("repo.private"),
	}
}
