package app

import (
	"groupup/src/models"
	"log"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err == nil {
		log.Println("sending delete cookie")
		logoutCookie := http.Cookie{
			Name:   "token",
			MaxAge: -1,
			Path:   "/",
		}
		http.SetCookie(w, &logoutCookie)
	}
	// TODO: Call something like "leaveGroup_helper or something that basically does what groupleave does.
	if models.UserExists(cookie.Value) {
		models.RemoveUser(cookie.Value)
	}
	http.Redirect(w, r, "/", http.StatusFound)
	return
}
