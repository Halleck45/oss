/**
 * This file is part of Open Source assets manager
 *
 * @license MIT
 * @copyright Jean-François Lépine
 */
package license

import (
	"os"
)

type License struct {
	Name string			`json:"name"`
	Identifier string	`json:"identifier"`
}

type Manifest struct {
	Assets []Asset		`json:"assets"`
}

type Asset struct {
	License License		`json:"license"`
	Description string	`json:"description"`
	File string			`json:"file"`
}

func (a *Asset) FileExists() (r bool){
	_, err := os.Stat(a.File)
	r = false
	if(err == nil) {
		r = true
	}
	return r
}
