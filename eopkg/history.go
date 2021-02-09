package eopkg

import "encoding/xml"

// History represents the history defined in package.
type History struct {
	XMLName xml.Name `xml:"History"`
	Update  Update
}

// Update represents the update information in the history.
type Update struct {
	XMLName xml.Name `xml:"Update"`
	Date    string
	Version string
	Comment string
	Name    string
	Email   string
}
