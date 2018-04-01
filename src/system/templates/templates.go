package templates

import (
	"html/template"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func LoadTemplateForApp(appURL string) (*template.Template, error) {
	allfiles, err := getAllFiles(appURL)
	if err != nil {
		return nil, err
	}

	return template.New(appURL).Delims("[[", "]]").ParseFiles(allfiles...)
}

func getAllFiles(appURL string) ([]string, error) {
	appindex := path.Join("./static", appURL, "index.html")
	globaltemplates, _ := filepath.Glob("./global/static/templates/*.html")
	apptemplates, _ := filepath.Glob("./static/portal/templates/*.html")

	_, err := os.Stat(appindex)
	if err != nil {
		return nil, err
	}

	allfiles := append(strings.Split(appindex, " "), append(globaltemplates, apptemplates...)...)
	return allfiles, nil
}
