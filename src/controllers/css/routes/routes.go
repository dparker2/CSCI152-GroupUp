package routes

import (
	"net/http"

	"groupup/pkg/types/routes"
	"groupup/src/controllers"
	CSSFileHandler "groupup/src/controllers/css/file"

	"github.com/go-xorm/xorm"
)

var db *xorm.Engine
var subrs map[string]routes.SubRoutePackage

func Init(c *controllers.MainController, DB *xorm.Engine) {
	db = DB
	//StatusHandler.Init(DB)

	c.RegisterSubRoute("/css",
		routes.Routes{
			routes.Route{"globalcss", "GET", "/global/{file:.*}", CSSFileHandler.Global},
			routes.Route{"appcss", "GET", "/{app}/{file:.*}", CSSFileHandler.App},
		},
		middleware,
	)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
