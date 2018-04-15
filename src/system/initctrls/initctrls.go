package initctrls

import (
	"groupup/src/controllers"
	approutes "groupup/src/controllers/app/routes"
	cssroutes "groupup/src/controllers/css/routes"
	fontroutes "groupup/src/controllers/fonts/routes"
	jsroutes "groupup/src/controllers/js/routes"
)

// InitCtrls initializes all other controllers
func InitCtrls(c *controllers.MainController) {
	//v1routes.Init(c)
	cssroutes.Init(c)
	jsroutes.Init(c)
	fontroutes.Init(c)
	approutes.Init(c)
}
