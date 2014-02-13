/**
 * Open Source assets manager
 *
 * @license MIT
 * @copyright Jean-François Lépine
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"io"
	"errors"
)


type Asset struct {
	License License		`json:"license"`
	Description string	`json:"description"`
	File string			`json:"file"`
}

type License struct {
	Name string			`json:"name"`
	Identifier string	`json:"identifier"`
}

type Manifest struct {
	Assets []Asset		`json:"assets"`
}

type SpdxLicense struct {
	Name string			`json:"name"`
	OsiApproved bool	`json:"osiApproved"`
}


/**
 * processor
 */
type Processor struct {
	filename string
	licenseFilename string
	Manif Manifest
}

/**
 * Initializing project
 */
func (c *Processor) init() (err error) {

	// local .oss file
	if _, err := os.Stat(c.filename); os.IsNotExist(err) {
		// first run: file doesn't exist yet
		c.Manif = Manifest{make([]Asset, 0, 100)}
		c.save()
	} else {
		err = c.load()
		if(err != nil) {
			return err
		}
	}

	// list of SPDX licenses
	err = c.update()
	return err
}

/**
 * Load current manifest
 */
func (c *Processor) load() (err error){
	// file exists ; load it
	file, err := os.Open(c.filename)
	if err != nil {
		fmt.Println("oss has not been initialized. Please run ./oss init")
		return err
	}

	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&c.Manif)
	return err
}


/**
 * Add asset in manifest
 */
func (c *Processor) add(a Asset) {
	c.Manif.Assets = append(c.Manif.Assets, a)
	c.save()
}

/**
 * Updates list of licenses
 * Big thanks to @sindresorhus for its job on https://github.com/sindresorhus/spdx-license-list
 */
func (c *Processor) update() (err error){

	fmt.Println("Downloading SPDX licenses list...")

	rawURL := "https://raw.githubusercontent.com/sindresorhus/spdx-license-list/master/spdx.json"
	fileName := c.licenseFilename
	file, err := os.Create(fileName)

	if err != nil {
		fmt.Printf("Cannot write to %s. Please check your permissions\n", fileName)
		return err
	}
	defer file.Close()

	check := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := check.Get(rawURL) // add a filter to check redirect

	if err != nil {
		fmt.Printf("Cannot download %s. Please check your connection\n", rawURL)
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	size, err := io.Copy(file, resp.Body)

	if err != nil {
		return err
	}

	fmt.Printf("%v bytes downloaded", size)
	fmt.Println("")

	return nil
}


/**
 * Factory license (only if SPDX license exists)
 */
func (c *Processor) factoryLicense(identifier string) (lic License, err error) {
	licenses := c.listSpdxLicenses()
	spdxLicense, ok := licenses[identifier]

	if(!ok) {
		return License{}, errors.New("Blank strings not accepted.")
	}

	lic = License{Identifier: identifier, Name: spdxLicense.Name}
	return lic, nil

}

/**
 * List Spdx licenses
 */
func (c *Processor) listSpdxLicenses() map[string]SpdxLicense {
	file, e := ioutil.ReadFile(c.licenseFilename)
	if e != nil {
		fmt.Println("Project has not been initialized. Please run './oss init'")
		os.Exit(1)
	}
	licenses := map[string]SpdxLicense{}
	json.Unmarshal(file, &licenses)
	return licenses
}

/**
 * Saving manifest
 */
func (c *Processor) save() {

	// json
	b, err := json.Marshal(c.Manif)
	if err != nil {
		panic(err)
	}

	// file
	err = ioutil.WriteFile(c.filename, b, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Manifest saved")
}



func main() {
	p := Processor{filename: "./.oss", licenseFilename: "./.oss-licenses.json"};

	var command string
	if(len(os.Args) < 2) {
		command = "help"
	} else {
		command = os.Args[1];
	}

	switch(command) {
		// Initialization
		// -------------------------
		case "init":
			err := p.init()
			if(err != nil) {
				fmt.Println("Project has not been initialized")
			}

		// Update list of licenses
		// -------------------------
		case "update":
			p.update()

		// List registered assets
		// -------------------------
		case "list":
			p.load()
			fmt.Printf("%d elements", len(p.Manif.Assets))
			fmt.Println("")
			for _, asset := range p.Manif.Assets {
				fmt.Printf("%-20s %-30s %-30s %s", asset.License.Identifier, asset.License.Name, asset.File, asset.Description)
				fmt.Println("")
			}

		// List available licenses
		// -------------------------
		case "list-licenses":
			licenses := p.listSpdxLicenses()
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
			p.load()
			for _, asset := range p.Manif.Assets {
				if(asset.File == file) {
					fmt.Printf("%-20s %-30s %-30s %s\n", asset.License.Identifier, asset.License.Name, asset.File, asset.Description)
					return
				}
			}
			fmt.Println("File is not registered")

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
			license, err := p.factoryLicense(lic)
			if(err != nil) {
				fmt.Printf("License %s is not registred in SPDX database\n", lic)
				return
			}

			p.load()
			a := Asset{License: license, File: file, Description: descr}
			p.add(a)

		// Display version
		// -------------------------
		case "version":
			fmt.Println(version)

		// Help
		// -------------------------
		case "help":
			fmt.Println("Tools for managing Open Source assets, by Jean-François Lépine <http://lepine.pro>")
			fmt.Println("")
			fmt.Println("Usage:")
			fmt.Println("./oss <command>")
			fmt.Printf("	%-20s %s", "init", "Initialize project\n")
			fmt.Printf("	%-20s %s", "add", "Add asset. Usage : add <license> <file> [<description>]\n")
			fmt.Printf("	%-20s %s", "show", "Display informations about file. Usage : show <file>\n")
			fmt.Printf("	%-20s %s", "update", "Update SPDX licenses list\n")
			fmt.Printf("	%-20s %s", "list", "List assets of the project\n")
			fmt.Printf("	%-20s %s", "list-licenses", "List SPDX licenses\n")
			fmt.Printf("	%-20s %s", "version", "Dislay current version\n")
			fmt.Printf("	%-20s %s", "help", "This help\n")
		default:
			fmt.Println("Command %s not found", command)
	}
}

var version string // will be initialized by build, or with run option
// for example: go run -ldflags "-X main.version xxx" oss.go
