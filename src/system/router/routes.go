package router

import (
	"log"
	"net/http"

	"groupup/pkg/types/routes"
	PortalHandler "groupup/src/controllers/portal"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Global middleware reached!")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes() (rts routes.Routes) {
	rts = routes.Routes{
		routes.Route{"Home", "GET", "/", PortalHandler.Index},
		routes.Route{"Portal", "GET", "/portal{extras:.*}", PortalHandler.Portal},
	}

	return
}
