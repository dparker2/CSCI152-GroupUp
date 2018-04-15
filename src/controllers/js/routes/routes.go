package routes

import (
	"net/http"

	"groupup/pkg/types/routes"
	"groupup/src/controllers"
	JSFileHandler "groupup/src/controllers/js/file"
)

var subrs map[string]routes.SubRoutePackage

func Init(c *controllers.MainController) {
	c.RegisterSubRoute("/js",
		routes.Routes{
			routes.Route{"globaljs", "GET", "/global/{file:.*}", JSFileHandler.Global},
			routes.Route{"supportjs", "GET", "/support/{file:.*}", JSFileHandler.Support},
			routes.Route{"appjs", "GET", "/{app}/{file:.*}", JSFileHandler.App},
		},
		middleware,
	)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
