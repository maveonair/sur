package eopkg

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"sort"
)

// ParseIndex parses the eopkg-index.xml and returns all packages.
func ParseIndex(indexFilePath string) ([]Package, error) {
	xmlFile, err := os.Open(indexFilePath)
	if err != nil {
		return nil, err
	}
	defer xmlFile.Close()

	content, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return nil, err
	}

	var packageIndex Packages
	xml.Unmarshal(content, &packageIndex)

	sort.Sort(ByName(packageIndex.Packages))
	return packageIndex.Packages, nil
}
