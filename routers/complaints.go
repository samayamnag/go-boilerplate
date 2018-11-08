package routers

import (
	"github.com/gorilla/mux"

	"github.com/janaonline/icmyc/app/controllers"
	"github.com/janaonline/icmyc/app/controllers/v1"
	"github.com/janaonline/icmyc/app/controllers/v2"
)

var complaintCtrl = &controllers.ComplaintController{}
var complaintCtrlV1 = &v1.ComplaintController{}
var complaintCtrlV2 = &v2.ComplaintController{}

func SetComplaintRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", complaintCtrl.Index).Methods("GET")
	router.HandleFunc("/complaints", complaintCtrl.Index).Methods("GET")
	
	return router
}

func SetComplaintV1Routes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", complaintCtrlV1.Index).Methods("GET", "POST")	
	return router
}

func SetComplaintV2Routes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", complaintCtrlV2.Index).Methods("GET")
	router.HandleFunc("/test", complaintCtrlV2.Index).Methods("GET")	
	return router
}

func SetComplaintV3Routes(router *mux.Router) *mux.Router {
	router.HandleFunc("/", complaintCtrl.Index).Methods("GET")	
	return router
}