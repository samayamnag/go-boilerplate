package routers

import (
	"github.com/gorilla/mux"
	_ "github.com/urfave/negroni"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetComplaintRoutes(router)
	router = SetUserRoutes(router)

	// Subrouters
	complaintRouterV1 := mux.NewRouter().PathPrefix("/v1/complaints").Subrouter().StrictSlash(true)
	complaintRouterV2 := mux.NewRouter().PathPrefix("/v2/complaints").Subrouter().StrictSlash(true)
	complaintRouterV3 := mux.NewRouter().PathPrefix("/v3/complaints").Subrouter().StrictSlash(true)

	// Subrouters
	userRouterV1 := mux.NewRouter().PathPrefix("/v1/users").Subrouter().StrictSlash(true)
	userRouterV2 := mux.NewRouter().PathPrefix("/v2/users").Subrouter().StrictSlash(true)

	// Set version specific routes
	complaintRouterV1 = SetComplaintV1Routes(complaintRouterV1)
	complaintRouterV2 = SetComplaintV2Routes(complaintRouterV2)
	complaintRouterV3 = SetComplaintV3Routes(complaintRouterV3)

	// Set version specific routes - Users
	userRouterV1 = SetUserV1Routes(userRouterV1)
	userRouterV2 = SetUserV2Routes(userRouterV2)

	// Link to main router - Complaints
	router.PathPrefix("/v1/complaints").Handler(complaintRouterV1)
	router.PathPrefix("/v2/complaints").Handler(complaintRouterV2)
	router.PathPrefix("/v3/complaints").Handler(complaintRouterV3)

	// Link to main router - Users
	router.PathPrefix("/v1/users").Handler(userRouterV1)
	router.PathPrefix("/v2/users").Handler(userRouterV2)
	
	return router
}
