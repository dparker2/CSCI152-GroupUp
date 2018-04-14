package portal

import (
	"encoding/json"
	"fmt"
	TemplateLoader "groupup/src/system/templates"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/portal", http.StatusFound)
}

func Portal(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	log.Println(r.URL.Path)
	//http.StripPrefix("/portal"+vars["extras"], http.FileServer(http.Dir("./static/portal/"))).ServeHTTP(w, r)
	tmpl, err := TemplateLoader.LoadTemplateForApp(r.URL.Path)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}

	tmpl.ExecuteTemplate(w, "portal", nil)
}

type test struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	log.Println(username)
	log.Println(password)
	if username == "developer" && password == "1234" {
		w.Header().Set("Content-Type", "application/json")
		p := test{Field1: username, Field2: password}
		// TODO: Pass a redirect path ("/app") to client, along with a random token. The client will then GET the path if it exists, putting the token in the header.
		// TODO: Make user struct with the token in it
		// TODO: Allow access to /app only with a valid token in header
		// TODO: Need a function in models like IsValidToken(token string) where the IP was the one associated with the token when it was first issues upon login. (maybe with a timeout?? like new time - old time < limit).
		// TODO: Also need models.AddUser(username, token)
		// TODO: When DB is good to go, need to do a verification here like models.VerifyLogin(username string, password string)
		json.NewEncoder(w).Encode(p)
	}
}
