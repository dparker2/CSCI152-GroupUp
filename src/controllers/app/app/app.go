package app

import (
	"fmt"
	TemplateLoader "groupup/src/system/templates"
	"log"
	"net/http"
)

func init() {
	setupAPI()
}

func App(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	tmpl, err := TemplateLoader.LoadTemplateForApp(r.URL.Path)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}

	tmpl.ExecuteTemplate(w, "app", nil)
}
