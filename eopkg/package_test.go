package eopkg

import (
	"sort"
	"testing"
)

func TestSortingPackages(t *testing.T) {
	packages := []Package{
		{
			Name: "strongswan",
		},
		{
			Name: "networkmanager-strongswan",
		},
		{
			Name: "font-jetbrains-mono-ttf",
		},
	}

	sort.Sort(ByName(packages))

	if packages[0].Name != "font-jetbrains-mono-ttf" {
		t.Errorf("Expected font-jetbrains-mono-ttf, actual: %s", packages[0].Name)
	}

	if packages[1].Name != "networkmanager-strongswan" {
		t.Errorf("Expected networkmanager-strongswan, actual: %s", packages[1].Name)
	}

	if packages[2].Name != "strongswan" {
		t.Errorf("Expected strongswan, actual: %s", packages[2].Name)
	}
}
