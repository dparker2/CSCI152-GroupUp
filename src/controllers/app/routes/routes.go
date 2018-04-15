package routes

import (
	"log"
	"net/http"

	"groupup/pkg/types/routes"
	"groupup/src/controllers"
	AppFileHandler "groupup/src/controllers/app/app"
	"groupup/src/models"
)

var subrs map[string]routes.SubRoutePackage

func Init(c *controllers.MainController) {
	c.RegisterSubRoute("/app",
		routes.Routes{
			routes.Route{"App", "GET", "/", AppFileHandler.App},
			routes.Route{"WebSocket", "GET", "/ws", AppFileHandler.WS},
			routes.Route{"Logout", "GET", "/logout", AppFileHandler.Logout},
		},
		middleware,
	)
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err == nil && models.UserExists(cookie.Value) {
			log.Println("good cookie")
			next.ServeHTTP(w, r)
		} else {
			if err != nil {
				log.Println(err.Error())
			}
			suicideCookie := http.Cookie{
				Name:   "token",
				MaxAge: -1,
			}
			http.SetCookie(w, &suicideCookie)
			log.Println("bad cookie")
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	})
}
