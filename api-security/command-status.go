// Copyright 2018. Akamai Technologies, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	api "github.com/akamai/AkamaiOPEN-edgegrid-golang/api-endpoints-v2"
	akamai "github.com/akamai/cli-common-golang"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var commandStatus cli.Command = cli.Command{
	Name:        "status",
	ArgsUsage:   "",
	Description: "Review the security status of an API Endpoint",
	HideHelp:    true,
	Action:      callStatus,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "endpoint",
			Usage: "The unique identifier for the endpoint.",
		},
		cli.StringFlag{
			Name:  "version",
			Usage: "The endpoint version number.",
		},
		cli.BoolFlag{
			Name:  "json",
			Usage: "Output JSON format",
		},
	},
}

func callStatus(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Getting Security Settings...",
		fmt.Sprintf("Getting Security Settings...... [%s]", color.GreenString("OK")),
	)

	version := c.Int("version")
	if version == 0 {
		version, err = api.GetLatestVersionNumber(c.Int("endpoint"), false)
		if err != nil {
			return output(c, nil, err)
		}
	}

	endpoint, err := api.GetVersion(c.Int("endpoint"), version)
	if err != nil {
		return output(c, endpoint, err)
	}

	e := &api.EndpointSecurity{
		APIEndPointID:              endpoint.APIEndPointID,
		APIEndPointName:            endpoint.APIEndPointName,
		SecurityScheme:             endpoint.SecurityScheme,
		AkamaiSecurityRestrictions: endpoint.AkamaiSecurityRestrictions,
	}

	return output(c, e, err)
}
