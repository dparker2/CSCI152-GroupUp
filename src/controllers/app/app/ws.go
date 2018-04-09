package app

import (
	"encoding/json"
	"fmt"
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

type wsMessage struct {
	Code     string `json:"code"`
	Groupid  string `json:"groupid"`
	Chat     string `json:"chat"`
	Username string `json:"username"`
}

func WS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		log.Println(messageType)
		log.Println(p)
		if err != nil {
			log.Println(err)
			return
		}

		var msg wsMessage
		err = json.Unmarshal(p, &msg)
		if err != nil {
			panic(err)
		}

		log.Println("JSON:")
		log.Println(msg)

		if msg.Code == "JOIN GROUP" {
			if !models.GroupExists(msg.Groupid) {
				models.AddGroup(msg.Groupid)
			}

			models.AddUserToGroup(msg.Username, conn, msg.Groupid)
		} else if msg.Code == "CHAT" {
			for _, c := range models.GetConnectionsInGroup(msg.Groupid) {
				c.WriteMessage(messageType, p)
			}
		}

		fmt.Printf("%+v", p)
	}
}
