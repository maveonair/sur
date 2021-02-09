package eopkg

import "encoding/xml"

// Packages represents the eopkg-index.xml root element.
type Packages struct {
	XMLName  xml.Name  `xml:"PISI"`
	Packages []Package `xml:"Package"`
}

// Package represents a package defined in epkg-index.xml.
type Package struct {
	XMLName     xml.Name `xml:"Package"`
	Name        string   `xml:"Name"`
	Summary     string   `xml:"Summary"`
	Description string   `xml:"Description"`
	License     string   `xml:"License"`
	History     History  `xml:"History"`
}

// ByName implements the sorting interface to sort a slice of packages.
type ByName []Package

func (a ByName) Len() int           { return len(a) }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
