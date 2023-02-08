// Copyright (c) 2019, Drone Plugins project authors
// Copyright (c) 2021, Robert Kaussow <mail@thegeeklab.de>

// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file.

package urfave

import (
	"context"
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/thegeeklab/drone-plugin-lib/v2/drone"
	"github.com/thegeeklab/drone-plugin-lib/v2/trace"
	"github.com/urfave/cli/v2"
)

const (
	NetDailerTimeout                 = 30 * time.Second
	HTTPTransportIdleTimeout         = 90 * time.Second
	HTTPTransportTLSHandshakeTimeout = 10 * time.Second
	HTTPTransportMaxIdleConns        = 100
)

// networkFlags has the cli.Flags for the drone.Network.
func networkFlags(category string) []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:     "transport.skip-verify",
			Usage:    "skip ssl verify",
			EnvVars:  []string{"PLUGIN_SKIP_VERIFY"},
			Category: category,
		},
	}
}

// NetworkFromContext creates a drone.Network from the cli.Context.
func NetworkFromContext(ctx *cli.Context) drone.Network {
	dialer := &net.Dialer{
		Timeout:   NetDailerTimeout,
		KeepAlive: NetDailerTimeout,
		DualStack: true,
	}

	transport := &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		DialContext:           dialer.DialContext,
		MaxIdleConns:          HTTPTransportMaxIdleConns,
		IdleConnTimeout:       HTTPTransportIdleTimeout,
		TLSHandshakeTimeout:   HTTPTransportTLSHandshakeTimeout,
		ExpectContinueTimeout: 1 * time.Second,
	}

	context := context.Background()
	skipVerify := ctx.Bool("transport.skip-verify")

	if skipVerify {
		logrus.Warning("ssl verification is turned off")

		transport.TLSClientConfig = &tls.Config{
			//nolint:gosec
			InsecureSkipVerify: true,
		}
	}

	if ctx.String("log-level") == logrus.TraceLevel.String() {
		context = trace.HTTP(context)
	}

	client := &http.Client{
		Transport: transport,
	}

	return drone.Network{
		Context:    context,
		SkipVerify: skipVerify,
		Client:     client,
	}
}
