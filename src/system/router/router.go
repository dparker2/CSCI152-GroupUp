package router

import (
	"groupup/pkg/types/routes"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"

	"groupup/src/controllers"
	"groupup/src/system/initctrls"
)

type Router struct {
	Router *mux.Router
}

func (r *Router) Init(db *xorm.Engine) {
	r.Router.Use(Middleware)

	baseRoutes := GetRoutes(db)
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	mainCtrl := new(controllers.MainController)

	initctrls.InitCtrls(mainCtrl, db)

	mainSubRts := mainCtrl.SubRoutePackages()
	for name, pack := range mainSubRts {
		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
	}
}

func (r *Router) AttachSubRouterWithMiddleware(path string, subroutes routes.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {
	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return
}

func NewRouter() (r Router) {
	r.Router = mux.NewRouter().StrictSlash(true)

	return
}

// func establishSubRoutes(srp routes.SubRoutePackage) {
// 	subRoutes := srp.GetRoutes(db)
// 	for name, pack := range subRoutes {
// 		r.AttachSubRouterWithMiddleware(name, pack.Routes, pack.Middleware)
// 	}
// }
