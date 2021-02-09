package eopkg

import "testing"

func TestParseIndex(t *testing.T) {
	packages, err := ParseIndex("../test/fixtures/eopkg-index.xml")
	if err != nil {
		t.Fatal(err)
	}

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

func TestParseIndexForHistory(t *testing.T) {
	packages, err := ParseIndex("../test/fixtures/eopkg-index.xml")
	if err != nil {
		t.Fatal(err)
	}

	if packages[0].History.Update.Version != "1.0.3" {
		t.Errorf("Expected 1.0.3, actual: %s", packages[0].History.Update.Version)
	}

	if packages[1].History.Update.Version != "1.4.5" {
		t.Errorf("Expected 1.4.5, actual: %s", packages[1].History.Update.Version)
	}

	if packages[2].History.Update.Version != "5.8.2" {
		t.Errorf("Expected 5.8.2, actual: %s", packages[2].History.Update.Version)
	}
}
