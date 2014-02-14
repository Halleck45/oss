/**
 * This file is part of Open Source assets manager
 *
 * @license MIT
 * @copyright Jean-François Lépine
 *
 *
 * This service allows to manage OSS entities
 */
package license

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"errors"
	"github.com/halleck45/oss/spdx"
)


/**
 * Service
 */
type Service struct {
	Filename string
	Manif Manifest
	SpdxService spdx.Service
}

/**
 * Initializing project
 */
func (c *Service) Init() (err error) {

	// local .oss file
	if _, err := os.Stat(c.Filename); os.IsNotExist(err) {
		// first run: file doesn't exist yet
		c.Manif = Manifest{make([]Asset, 0, 100)}
		c.Save()
	} else {
		err = c.Load()
		if(err != nil) {
			return err
		}
	}

	// spdx licenses
	return c.SpdxService.Update()
}

/**
 * Load current manifest
 */
func (c *Service) Load() (err error){
	// file exists ; load it
	file, err := os.Open(c.Filename)
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
func (c *Service) Add(a Asset) {
	c.Manif.Assets = append(c.Manif.Assets, a)
	c.Save()
}

/**
 * Get asset by its filename
 *
 * @param string
 * @return Asset
 */
func (c *Service) Get(file string) (asset Asset, err error) {
	for _, asset := range c.Manif.Assets {
		if(asset.File == file) {
			return asset, nil
		}
	}
	return Asset{}, errors.New("Asset not found")
}

/**
 * Removes Asset
 *
 * @param Asset
 */
func (c *Service) Remove(asset Asset) {
	for i, a := range c.Manif.Assets {
		// file was found
		if a.File == asset.File {
			// remove item
			c.Manif.Assets = c.Manif.Assets[:i+copy(c.Manif.Assets[i:], c.Manif.Assets[i+1:])]
			c.Save()
			return
		}
	}
}

/**
 * Saving manifest
 */
func (c *Service) Save() {

	// json
	b, err := json.Marshal(c.Manif)
	if err != nil {
		panic(err)
	}

	// file
	err = ioutil.WriteFile(c.Filename, b, 0644)
	if err != nil {
		panic(err)
	}
}
