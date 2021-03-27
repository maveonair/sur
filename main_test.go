package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestBaseURLInTemplate(t *testing.T) {
	baseURL := "http://localhost:5000"
	err := os.Setenv("SUR_BASE_URL", baseURL)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(packagesIndex)
	handler.ServeHTTP(res, req)

	if res.Code != 200 {
		t.Errorf("Expected HTTP status 200, found %d", res.Code)
	}

	body := res.Body.String()
	if !strings.Contains(body, baseURL) {
		t.Errorf("expected %s to be the base URL", baseURL)
	}
}

func TestFsPackages(t *testing.T) {
	req, err := http.NewRequest("GET", "/packages/eopkg-index.xml", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler := http.HandlerFunc(packagesIndex)
	handler.ServeHTTP(res, req)

	if res.Code != 200 {
		t.Errorf("Expected HTTP status 200, found %d", res.Code)
	}
}
