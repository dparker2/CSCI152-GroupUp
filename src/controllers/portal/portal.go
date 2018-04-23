package portal

import (
	"encoding/json"
	"fmt"
	"groupup/src/models"
	TemplateLoader "groupup/src/system/templates"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/portal", http.StatusFound)
}

func Portal(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("token")
	if err == nil { // Redirect to app if token exists (already logged in)
		http.Redirect(w, r, "/app", http.StatusFound)
		return
	}

	tmpl, err := TemplateLoader.LoadTemplateForApp(r.URL.Path)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "500 Internal Server Error", 500)
		return
	}

	tmpl.ExecuteTemplate(w, "portal", nil)
}

type loginResponse struct {
	RedirectPath string `json:"redirect-path"`
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
	if username == "developer" && password == "1234" { // TODO: Replace with models.VerifyLogin(username, password)
		// Send back a json object with where the client should go to
		// and a token to include in the header for authentication
		w.Header().Set("Content-Type", "application/json")
		// First need to make the user
		userToken := models.NewUser(username)
		userCookie := http.Cookie{Name: "token", Value: userToken, Path: "/", MaxAge: 86400}
		http.SetCookie(w, &userCookie)
		p := loginResponse{ // Ajax POST, so need to redirect on the client side
			RedirectPath: "/app", // At some point we could probably remember their last page and put them there?
		}
		// TODO: The client will then GET the path if it exists, putting the token in the header.
		// TODO: Allow access to /app only with a valid token in header
		// TODO: Need a function in models like IsValidToken(token string) where the IP was the one associated with the token when it was first issues upon login. (maybe with a timeout?? like new time - old time < limit).
		// TODO: Also need models.AddUser(username, token)
		// TODO: When DB is good to go, need to do a verification here like models.VerifyLogin(username string, password string)
		json.NewEncoder(w).Encode(p)
	}
}
