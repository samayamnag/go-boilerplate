package v1

import (
	"net/http"
	"github.com/janaonline/icmyc/app/services"
	"github.com/janaonline/icmyc/app/models"
	"github.com/janaonline/icmyc/app/handlers"
	"github.com/janaonline/icmyc/app/util"
	"github.com/janaonline/icmyc/app/requests"

	"github.com/gorilla/mux"
)

type UsersController struct{
	Id 			string	`json:"id,omitempty"`
	Email 		string	`json:"email,omitempty"`
	Timestamp	string	`json:"created_at,omitempty"`
}

func (c *UsersController) Store(w http.ResponseWriter, r *http.Request) {

	e := requests.CreateUserRequest(r)
	err := map[string]interface{}{"errors": e}
	errSize := len(e)
	res := handlers.Response{ResponseWriter: w}
	if errSize == 0 {
		var user models.User
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")

		// Store into DB
		user = services.InsertUser(user)

		// Send response
		res := handlers.Response{ResponseWriter: w}
		res.SendCreated(UsersController{
			Id: user.ID.Hex(),
			Email: user.Email,
			Timestamp: util.FormatMongoDate(user.CreatedAt),
		})
	} else {
		res.SendUnprocessableEntity(err) 
	}
}

func (c *UsersController) Index(w http.ResponseWriter, r *http.Request) {
	dbUsers := services.GetAllUsers()
	user := UsersController{}
	users := []UsersController{}

	for _, u := range dbUsers{
		user.Id = u.ID.Hex()
		user.Email = u.Email
		user.Timestamp = util.FormatMongoDate(u.CreatedAt)
		
		users = append(users, user)
	}

	res := handlers.Response{ResponseWriter: w}

	res.SendOK(users)
}

func (c *UsersController) Show(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	dbUser := services.FindById(params["id"])

	res := handlers.Response{ResponseWriter: w}
	res.SendOK(UsersController{
		Id: dbUser.ID.Hex(),
		Email: dbUser.Email,
		Timestamp: util.FormatMongoDate(dbUser.CreatedAt),
	})
}