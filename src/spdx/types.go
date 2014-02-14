/**
 * This file is part of Open Source assets manager
 *
 * @license MIT
 * @copyright Jean-François Lépine
 */
package spdx


type License struct {
	Name string			`json:"name"`
	Identifier string	`json:"identifier"`
}

type SpdxLicense struct {
	Name string			`json:"name"`
	OsiApproved bool	`json:"osiApproved"`
}
