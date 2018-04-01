package routes

import (
	"net/http"
)

type SubRouteInterface interface {
	GetRoutes() (SubRoute map[string]SubRoutePackage)
}

type Routes []Route

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type SubRoutePackage struct {
	Routes     Routes
	Middleware func(next http.Handler) http.Handler
}
