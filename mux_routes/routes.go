package mux_routes

import (
	"net/http"

	"github.com/janaonline/icmyc/app/controllers"
	"github.com/janaonline/icmyc/app/controllers/v1"
	"github.com/janaonline/icmyc/app/controllers/v2"
)

var complaintCtrl = &controllers.ComplaintController{}
var complaintCtrlV1 = &v1.ComplaintController{}
var complaintCtrlV2 = &v2.ComplaintController{}

type Route struct {
	Name		string
	Method		string
	Pattern		string
	HandlerFunc	http.HandlerFunc 
}

type Routes []Route

var defaultRouters = Routes {
	Route{
		"GetComplaints",
		"GET",
		"/",
		complaintCtrl.Index,
	},
	Route{
		"GetComplaints",
		"GET",
		"/test",
		complaintCtrl.Index,
	},
}

var routesV1 = Routes {
	Route{
		"GetComplaints",
		"GET",
		"/",
		complaintCtrlV1.Index,
	},
}

var routesV2 = Routes {
	Route{
		"GetComplaints",
		"GET",
		"/",
		complaintCtrlV2.Index,
	},
}

var routesV3 = Routes {
	Route{
		"GetComplaints",
		"GET",
		"/",
		complaintCtrl.Index,
	},
	Route{
		"GetComplaints",
		"GET",
		"/test",
		complaintCtrl.Index,
	},
}

