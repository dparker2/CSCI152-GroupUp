package routes

import (
	"net/http"

	"groupup/pkg/types/routes"
	"groupup/src/controllers"
	FontsFileHandler "groupup/src/controllers/fonts/file"
)

var subrs map[string]routes.SubRoutePackage

func Init(c *controllers.MainController) {
	c.RegisterSubRoute("/fonts",
		routes.Routes{
			routes.Route{"globalfonts", "GET", "/{file:.*}", FontsFileHandler.Global},
		},
		middleware,
	)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
