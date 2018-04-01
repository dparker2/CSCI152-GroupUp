package initctrls

import (
	"groupup/src/controllers"
	approutes "groupup/src/controllers/app/routes"
	cssroutes "groupup/src/controllers/css/routes"
	fontroutes "groupup/src/controllers/fonts/routes"
	jsroutes "groupup/src/controllers/js/routes"
	v1routes "groupup/src/controllers/v1/routes"

	"github.com/go-xorm/xorm"
)

// InitCtrls initializes all other controllers
func InitCtrls(c *controllers.MainController, DB *xorm.Engine) {
	v1routes.Init(c, DB)
	cssroutes.Init(c, DB)
	jsroutes.Init(c, DB)
	fontroutes.Init(c, DB)
	approutes.Init(c, DB)
}
