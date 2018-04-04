package routes

import (
	"log"
	"net/http"

	"groupup/pkg/types/routes"
	"groupup/src/controllers"
	AppFileHandler "groupup/src/controllers/app/app"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine
var subrs map[string]routes.SubRoutePackage

func Init(c *controllers.MainController, DB *xorm.Engine) {
	db = DB
	//StatusHandler.Init(DB)

	c.RegisterSubRoute("/app",
		routes.Routes{
			routes.Route{"App", "GET", "/", AppFileHandler.App},
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
