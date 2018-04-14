package routes

import (
	"log"
	"net/http"

	"groupup/pkg/types/routes"
	"groupup/src/controllers"
	AppFileHandler "groupup/src/controllers/app/app"
)

var subrs map[string]routes.SubRoutePackage

func Init(c *controllers.MainController) {
	c.RegisterSubRoute("/app",
		routes.Routes{
			routes.Route{"App", "GET", "/", AppFileHandler.App},
			routes.Route{"WebSocket", "GET", "/ws", AppFileHandler.WS},
		},
		middleware,
	)
}

func middleware(next http.Handler) http.Handler {
	log.Println("TODO: Authentication... verify x-app-token header to allow access.")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
