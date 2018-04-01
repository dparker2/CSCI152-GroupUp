package file

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Global(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := "./global/static/js/" + vars["file"]
	log.Println(filename)

	serveIfExists(w, r, filename)
}

func Support(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := "./node_modules/" + vars["file"]
	log.Println(filename)

	serveIfExists(w, r, filename)
}

func App(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := "./static/" + vars["app"] + "/js/" + vars["file"]
	log.Println(filename)

	serveIfExists(w, r, filename)
}

func serveIfExists(w http.ResponseWriter, r *http.Request, filename string) {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		log.Println("it exists...")
		http.ServeFile(w, r, filename)
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
