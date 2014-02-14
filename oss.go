/**
 * This file is part of Open Source assets manager
 *
 * @license MIT
 * @copyright Jean-François Lépine
 */
package main

import (
	"os"
	"github.com/halleck45/oss/license"
	"github.com/halleck45/oss/spdx"
)

func main() {

	var command string
	if(len(os.Args) < 2) {
		command = "help"
	} else {
		command = os.Args[1];
	}

	spdxService := spdx.Service{LicenseFilename: "./.oss-licenses.json"}
	service := license.Service{SpdxService: spdxService, Filename: "./.oss"}
	app := license.Application{Service: service, Version: version}
	app.Run(command)

}

var version string // will be initialized by build, or with run option
// for example: go run -ldflags "-X main.version xxx" oss.go
