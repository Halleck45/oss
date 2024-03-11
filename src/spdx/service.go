/**
 * This file is part of Open Source assets manager
 *
 * @license MIT
 * @copyright Jean-François Lépine
 *
 *
 * service allows to manage SPDX entities
 */
package spdx

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"io"
	"errors"
	"github.com/dotcypress/phonetics"
	"regexp"
)


/**
 * SdpxService
 */
type Service struct {
	LicenseFilename string
}

/**
 * Updates list of licenses
 * Big thanks to @sindresorhus for its job on https://github.com/sindresorhus/spdx-license-list
 */
func (c *Service) Update() (err error){

	fmt.Println("Downloading SPDX licenses list...")

	rawURL := "https://raw.githubusercontent.com/sindresorhus/spdx-license-list/master/spdx.json"
	Filename := c.LicenseFilename
	file, err := os.Create(Filename)

	if err != nil {
		fmt.Printf("Cannot write to %s. Please check your permissions\n", Filename)
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
func (c *Service) Get(identifier string) (lic License, err error) {
	licenses := c.All()
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
func (c *Service) All() map[string]SpdxLicense {
	file, e := ioutil.ReadFile(c.LicenseFilename)
	if e != nil {
		fmt.Println("Project has not been initialized. Please run './oss init'")
		os.Exit(1)
	}
	licenses := map[string]SpdxLicense{}
	json.Unmarshal(file, &licenses)
	return licenses
}

/**
 * Search term by metaphon
 */
func (c *Service) SearchAround(term string) (License, error) {

	max := 0
	var bestMatch string
	for identifier, _:= range c.All() {
		proximity := phonetics.DifferenceSoundex(term, identifier)
		if(proximity > max) {
			bestMatch = identifier
			max = proximity
		}
	}

	// factory license
	license, _ := c.Get(bestMatch)
	return license, nil

}

/**
 * Search term matching
 */
func (c *Service) SearchMatching(term string) (map[string]License) {
	licenses := map[string]License{}
	r, _ := regexp.Compile(fmt.Sprintf("(%s+)", term))
	for identifier, spdxLicense:= range c.All() {
		if r.MatchString(identifier) || r.MatchString(spdxLicense.Name) {
			license, _ := c.Get(identifier)

			// replace matched part
			replacement := fmt.Sprintf("%s%s%s", "\x1b[32m", term, "\x1b[0m")
			license.Identifier = r.ReplaceAllString(license.Identifier, replacement)
			license.Name = r.ReplaceAllString(license.Name, replacement)

			licenses[identifier] = license

		}
	}
	return licenses
}
