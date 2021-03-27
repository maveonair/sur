package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"git.sr.ht/~maveonair/sur/eopkg"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// PageData represents the context for the templates.
type PageData struct {
	Packages []eopkg.Package
	BaseURL  string
}

func main() {
	fsPackages := http.FileServer(http.Dir(getPackagesDirectory()))

	router := mux.NewRouter()
	router.PathPrefix("/packages").Handler(fsPackages)
	router.HandleFunc("/", packagesIndex).Methods("GET")

	handler := handlers.LoggingHandler(os.Stdout, router)
	server := &http.Server{
		Handler:      handlers.CompressHandler(handler),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Info("Listening on port 8080...")
	log.Fatal(server.ListenAndServe())
}

func packagesIndex(w http.ResponseWriter, r *http.Request) {
	indexFilePath := filepath.Join(getPackagesDirectory(), "eopkg-index.xml")
	packages, err := eopkg.ParseIndex(indexFilePath)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	baseURL := getBaseURL(r)

	pageData := PageData{
		Packages: packages,
		BaseURL:  baseURL,
	}

	tmpl := template.Must(template.ParseFiles("tmpl/index.html", "tmpl/layout.html"))
	tmpl.ExecuteTemplate(w, "layout", pageData)
}

func getBaseURL(r *http.Request) string {
	baseURL := os.Getenv("SUR_BASE_URL")
	if baseURL != "" {
		return baseURL
	}

	return fmt.Sprintf("http://%s", r.Host)
}

func getPackagesDirectory() string {
	packageDir := os.Getenv("SUR_PACKAGES_DIR")
	if packageDir == "" {
		log.Fatal("SUR_PACKAGES_DIR not defined")
	}

	return packageDir
}
