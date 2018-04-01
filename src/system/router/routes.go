package router

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"

	"groupup/pkg/types/routes"
	PortalHandler "groupup/src/controllers/portal"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Global middleware reached!")
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) (rts routes.Routes) {
	PortalHandler.Init(db)

	rts = routes.Routes{
		routes.Route{"Home", "GET", "/", PortalHandler.Index},
		routes.Route{"Portal", "GET", "/portal{extras:.*}", PortalHandler.Portal},
	}

	return
}
