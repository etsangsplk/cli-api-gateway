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

	api "github.com/akamai/AkamaiOPEN-edgegrid-golang/apikey-manager-v1"
	akamai "github.com/akamai/cli-common-golang"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

var commandAclAllow cli.Command = cli.Command{
	Name:        "acl-allow",
	ArgsUsage:   "",
	Description: "This operation adds/allows an endpoint on the key collection ACL",
	HideHelp:    true,
	Action:      callAclAllow,
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:  "collection",
			Usage: "The collection ID to modify.",
		},
		cli.IntSliceFlag{
			Name:  "endpoint",
			Usage: "The endpoint ID to add to the ACL, multiples allowed",
		},
		cli.IntSliceFlag{
			Name:  "resource",
			Usage: "The resource ID to add to the ACL, multiples allowed",
		},
	},
}

func callAclAllow(c *cli.Context) error {
	err := initConfig(c)
	if err != nil {
		return cli.NewExitError(color.RedString(err.Error()), 1)
	}

	akamai.StartSpinner(
		"Updating key collection ACL...",
		fmt.Sprintf("Updating key collection ACL...... [%s]", color.GreenString("OK")),
	)

	var acl []string

	for _, e := range c.IntSlice("endpoint") {
		acl = append(acl, fmt.Sprintf("ENDPOINT-%d", e))
	}

	for _, r := range c.IntSlice("resource") {
		acl = append(acl, fmt.Sprintf("RESOURCE-%d", r))
	}

	collection, err := api.CollectionAclAllow(c.Int("collection"), acl)

	return output(c, collection, err)
}