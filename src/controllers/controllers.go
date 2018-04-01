package controllers

import (
	"groupup/pkg/types/routes"
	"net/http"
)

// MainController holds a map of all sub route packages registered.
type MainController struct {
	subPackages map[string]routes.SubRoutePackage
}

// SubRoutePackages returns the map of all sub route packages registered.
func (c *MainController) SubRoutePackages() map[string]routes.SubRoutePackage {
	return c.subPackages
}

// RegisterSubRoute adds routes.Routes rs and func(next http.Handler) http.Handler mw to a map with key path
func (c *MainController) RegisterSubRoute(path string, rs routes.Routes, mw func(next http.Handler) http.Handler) {
	if c.subPackages == nil {
		c.subPackages = make(map[string]routes.SubRoutePackage)
	}
	c.subPackages[path] = routes.SubRoutePackage{
		Routes:     rs,
		Middleware: mw,
	}

	return
}
