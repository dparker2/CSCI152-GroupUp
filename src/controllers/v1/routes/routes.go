package routes

import (
	"log"
	"net/http"

	"groupup/pkg/types/routes"
	"groupup/src/controllers"
	StatusHandler "groupup/src/controllers/v1/status"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine

func Init(c *controllers.MainController, DB *xorm.Engine) {
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
