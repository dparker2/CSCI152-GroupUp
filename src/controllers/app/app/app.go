package app

import (
	"fmt"
	TemplateLoader "groupup/src/system/templates"
	"log"
	"net/http"
)

func App(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	log.Println(r.URL.Path)
	//http.StripPrefix("/portal"+vars["extras"], http.FileServer(http.Dir("./static/portal/"))).ServeHTTP(w, r)
	tmpl, err := TemplateLoader.LoadTemplateForApp(r.URL.Path)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}

	tmpl.ExecuteTemplate(w, "app", nil)
}
