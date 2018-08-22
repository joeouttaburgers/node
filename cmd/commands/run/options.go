/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package run

import (
	"flag"

	"github.com/mysterium/node/cmd"
	"github.com/mysterium/node/core/node"
)

// CommandOptions describes options which are required to start Command
type CommandOptions struct {
	CLI               bool
	Version           bool
	LicenseWarranty   bool
	LicenseConditions bool

	NodeOptions node.NodeOptions
}

// ParseArguments parses CLI flags and adds to CommandOptions structure
func ParseArguments(args []string) (options CommandOptions, err error) {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)

	err = cmd.ParseFromCmdArgs(flags, &options.NodeOptions.Directories)
	if err != nil {
		return
	}

	flags.StringVar(
		&options.NodeOptions.OpenvpnBinary,
		"openvpn.binary",
		"openvpn", //search in $PATH by default,
		"openvpn binary to use for Open VPN connections",
	)

	flags.StringVar(
		&options.NodeOptions.TequilapiAddress,
		"tequilapi.address",
		"127.0.0.1",
		"IP address of interface to listen for incoming connections",
	)
	flags.IntVar(
		&options.NodeOptions.TequilapiPort,
		"tequilapi.port",
		4050,
		"Port for listening incoming api requests",
	)

	flags.BoolVar(
		&options.CLI,
		"cli",
		false,
		"Run an interactive CLI based Mysterium UI",
	)
	flags.BoolVar(
		&options.Version,
		"version",
		false,
		"Show version",
	)
	flags.BoolVar(
		&options.LicenseWarranty,
		"license.warranty",
		false,
		"Show warranty",
	)
	flags.BoolVar(
		&options.LicenseConditions,
		"license.conditions",
		false,
		"Show conditions",
	)

	flags.StringVar(
		&options.NodeOptions.IpifyUrl,
		"ipify-url",
		"https://api.ipify.org/",
		"Address (URL form) of ipify service",
	)

	flags.StringVar(
		&options.NodeOptions.LocationDatabase,
		"location.database",
		"GeoLite2-Country.mmdb",
		"Service location autodetect database of GeoLite2 format e.g. http://dev.maxmind.com/geoip/geoip2/geolite2/",
	)

	cmd.ParseNetworkOptions(flags, &options.NodeOptions.NetworkOptions)

	err = flags.Parse(args[1:])
	if err != nil {
		return
	}

	return options, err
}
