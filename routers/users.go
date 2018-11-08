package routers

import (
	"github.com/gorilla/mux"

	"github.com/janaonline/icmyc/app/controllers"
	"github.com/janaonline/icmyc/app/controllers/v1"
)

var usersCtrl = &controllers.UsersController{}
var usersCtrlV1 = &v1.UsersController{}

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users", usersCtrl.Store).Methods("POST")
	router.HandleFunc("/users", usersCtrl.Index).Methods("GET")
	router.HandleFunc("/users/{id}", usersCtrl.Show).Methods("GET")
	
	return router
}

func SetUserV1Routes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", usersCtrlV1.Store).Methods("POST")
	router.HandleFunc("/", usersCtrlV1.Index).Methods("GET")
	router.HandleFunc("/{id}", usersCtrlV1.Show).Methods("GET")
	
	return router
}

func SetUserV2Routes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", usersCtrl.Store).Methods("POST")
	router.HandleFunc("/", usersCtrl.Index).Methods("GET")
	router.HandleFunc("/{id}", usersCtrl.Show).Methods("GET")
	
	return router
}