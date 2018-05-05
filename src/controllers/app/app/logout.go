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

	userToken := cookie.Value

	letFollowersKnow(userToken, &wsMessage{
		Code:     "app/friends/offline",
		Username: models.GetUsername(userToken),
	})
	// TODO: Call something like "leaveGroup_helper or something that basically does what groupleave does.
	if models.UserExists(userToken) {
		models.RemoveUser(userToken)
	}
	http.Redirect(w, r, "/", http.StatusFound)
	return
}

func letFollowersKnow(token string, msg *wsMessage) (err error) {
	onlineFollowers, err := models.GetOnlineFollowers(token)
	if err != nil {
		return
	}
	for _, friendName := range onlineFollowers {
		if models.UserExistsByUsername(friendName) {
			conn := models.GetConnectionByUsername(friendName)
			err = conn.WriteJSON(msg)
		}
	}
	return
}
