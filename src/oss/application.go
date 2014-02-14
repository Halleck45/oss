/**
 * This file is part of Open Source assets manager
 *
 * @license MIT
 * @copyright Jean-François Lépine
 *
 * Represents the application layer
 */
package oss

import (
	"fmt"
	"os"
)

/**
 * Application
 */
type Application struct {
	Service Service
	Version string
}

/**
 * Initializing project
 */
func (app *Application) Run(command string) {

	switch(command) {
		// Initialization
		// -------------------------
		case "init":
			err := app.Service.Init()
			if(err != nil) {
				fmt.Println("Project has not been initialized")
			}

		// Update list of licenses
		// -------------------------
		case "update":
			app.Service.Spdx.Update()

		// List registered assets
		// -------------------------
		case "status":
			app.Service.Load()
			fmt.Printf("%d elements", len(app.Service.Manif.Assets))
			fmt.Println("")
			for _, asset := range app.Service.Manif.Assets {

				// check file
//				fileExists := fmt.Sprintf("%s %s %s", "\x1b[32m", "✔ file exists", "\x1b[0m")
//				if(!asset.FileExists()) {
//					fileExists = fmt.Sprintf("%s %s %s", "\x1b[31m", "✘ not found", "\x1b[0m")
//				}
				fileExists := fmt.Sprintf("%s%s", "\x1b[32m", "✔")
				if(!asset.FileExists()) {
					fileExists = fmt.Sprintf("%s%s", "\x1b[31m", "✘")
				}

				fmt.Printf("     %-15s %-5s %-50s \x1b[0m %s %s",
					asset.License.Identifier,
					fileExists,
					asset.File,
					asset.License.Name,
					asset.Description)
				fmt.Println("")
			}

		// List available licenses
		// -------------------------
		case "licenses":
			licenses := app.Service.Spdx.All()
			for identifier, license := range licenses {
				fmt.Printf("%-30s %s", identifier, license.Name)
				fmt.Println("")
			}

		//
		// Display information about one file
		// -------------------------
		case "show":
			if(len(os.Args) < 3) {
				fmt.Println("Parameters are missing")
				fmt.Println("Usage: ./oss show <file>")
				return;
			}
			file := os.Args[2]
			app.Service.Load()
			asset, err := app.Service.Get(file)
			if(err != nil) {
				fmt.Println("File is not registered")
				return
			}
			fmt.Printf("%-20s %-30s %-30s %s\n", asset.License.Identifier, asset.License.Name, asset.File, asset.Description)

		// Registering asset
		// -------------------------
		case "add":
			if(len(os.Args) < 3) {
				fmt.Println("Parameters are missing")
				fmt.Println("Usage: ./oss add <license> <file> [<description>]")
				return;
			}
			lic := os.Args[2];
			file := os.Args[3];
			var descr string
			if(len(os.Args) > 4) {
				descr = os.Args[4];
			}

			// check if license is valid
			spdxlicense, err := app.Service.Spdx.Get(lic)
			if(err != nil) {
				fmt.Printf("License %s is not registred in SPDX database\n", lic)
				return
			}
			app.Service.Load()
			license := License{Identifier: lic, Name: spdxlicense.Name}
			asset := Asset{License: license, File: file, Description: descr}
			app.Service.Add(asset)

		// Display version
		// -------------------------
		case "rm":
			if(len(os.Args) < 3) {
				fmt.Println("Parameters are missing")
				fmt.Println("Usage: ./oss show <file>")
				return;
			}
			file := os.Args[2]
			app.Service.Load()

			// try to get asset
			asset, err := app.Service.Get(file)
			if(err != nil) {
				fmt.Println("File is not registered")
				return
			}

			// remove
			app.Service.Remove(asset)
			fmt.Printf("%s has been removed from licenses list\n", file)
			return


		// Display version
		// -------------------------
		case "version":
			fmt.Println(app.Version)

		// Help
		// -------------------------
		case "help":
			fmt.Println("Tools for managing Open Source assets, by Jean-François Lépine <http://lepine.pro>")
			fmt.Println("")
			fmt.Println("Usage:")
			fmt.Println("./oss <command>")
			fmt.Printf("	%-20s %s", "init", "Initialize project\n")
			fmt.Printf("	%-20s %s", "add", "Add asset. Usage : add <license> <file> [<description>]\n")
			fmt.Printf("	%-20s %s", "rm", "Remove asset. Usage : rm <file>\n")
			fmt.Printf("	%-20s %s", "show", "Display informations about file. Usage : show <file>\n")
			fmt.Printf("	%-20s %s", "update", "Update SPDX licenses list\n")
			fmt.Printf("	%-20s %s", "status", "List assets of the project\n")
			fmt.Printf("	%-20s %s", "licenses", "List SPDX licenses\n")
			fmt.Printf("	%-20s %s", "version", "Dislay current version\n")
			fmt.Printf("	%-20s %s", "help", "This help\n")
		default:
			fmt.Println("Command %s not found", command)
	}
}
