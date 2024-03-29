package app

import (
	"log"
	"net/http"

	"groupup/src/models"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		log.Println("CHECKING ORIGIN")
		return true
	},
}

func WS(w http.ResponseWriter, r *http.Request) {
	// Get the token cookie
	cookie, err := r.Cookie("token")
	if err != nil {
		return
	}
	token := cookie.Value

	// Set user offline after disconnection
	defer models.SetUserStatus(token, 0)

	// Upgrade the connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Set user online
	models.SetUserStatus(token, 1)
	models.SetUserConn(token, conn)

	for {
		// Decode JSON received, wsMessage defines the supported parameters
		var msg wsMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			// TODO: Put a function here that is like "cleanupuser" where it removes them from the group theyre active in and puts them offline, etc
			//i.e. our disconnect and removing them from everything.
			letFollowersKnow(token, &wsMessage{
				Code:     "app/friends/offline",
				Username: models.GetUsername(token),
			})
			return
		}

		log.Println(msg)

		// Call the function the code corresponds to the received code
		if f, exists := wsAPI[msg.Code]; exists {
			err := f(wsAPIstruct{
				UserToken: token,
				Msg:       &msg,
			})
			if err != nil {
				log.Println(err.Error())
			}
		}
	}
}

func writeJSONToGroup(grpName string, msgJSON *wsMessage) {
	for _, c := range models.GetConnectionsInGroup(grpName) {
		err := c.WriteJSON(msgJSON)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func writeJSONToOthersInGroup(grpName string, token string, msgJSON *wsMessage) {
	for _, c := range models.GetOtherConnectionsInGroup(token, grpName) {
		err := c.WriteJSON(msgJSON)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
