package mux_routes

import (
	"net/http"
	"github.com/gorilla/mux"
)

func DefaultRouter(r *mux.Router) *mux.Router {
	for _, route := range defaultRouters {
		var handler http.Handler

		handler = route.HandlerFunc
		
		r.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)		
	}
	
	return r
}

func AddV1Routes(r *mux.Router) *mux.Router {
	for _, route := range routesV1 {
		var handler http.Handler

		handler = route.HandlerFunc
		
		r.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)		
	}
	
	return r
}

func AddV2Routes(r *mux.Router) *mux.Router {
	for _, route := range routesV2 {
		var handler http.Handler

		handler = route.HandlerFunc
		
		r.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)		
	}
	
	return r
}

func AddV3Routes(r *mux.Router) *mux.Router {
	for _, route := range routesV3 {
		var handler http.Handler

		handler = route.HandlerFunc
		
		r.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(handler)		
	}
	
	return r
}