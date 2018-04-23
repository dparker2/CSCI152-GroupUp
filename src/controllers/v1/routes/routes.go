package routes

import (
	"database/sql"
	"log"
	"net/http"

	"groupup/pkg/types/routes"
	"groupup/src/controllers"
	StatusHandler "groupup/src/controllers/v1/status"
)

var db *sql.DB

func Init(c *controllers.MainController, DB *sql.DB) {
	db = DB
	StatusHandler.Init(DB)

	c.RegisterSubRoute("/v1",
		routes.Routes{
			routes.Route{"Status", "GET", "/status", StatusHandler.Index},
		},
		middleware,
	)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("v1 middleware reached!")
		token := r.Header.Get("X-App-Token")
		if len(token) < 1 {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
