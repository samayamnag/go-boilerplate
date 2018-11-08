package v1

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/janaonline/icmyc/app/handlers"
	"github.com/janaonline/icmyc/app/requests"
)

type ComplaintController struct{
	Id int `json:"id,omitempty"`
	Desc string `json:"body,omitempty"`

}

func (c *ComplaintController) Index(w http.ResponseWriter, r *http.Request)  {
	a := requests.Address{
		Street: r.FormValue("street"),
		City:   r.FormValue("city"),
		State: r.FormValue("state"),
		Zip:    r.FormValue("zip"),
	}

	err := a.Validate()
	b, _ := json.Marshal(err)
	fmt.Println(err)

	res := handlers.Response{ResponseWriter: w}
	if err == nil {
		res.SendOK(ComplaintController{Id: 1, Desc:"Hello V1"})
	} else {
		res.SendValidationErrors(b) 
	}

	/*e := requests.CreateComplaintRequest(r)
	err := map[string]interface{}{"errors": e}
	errSize := len(e)
	res := handlers.Response{ResponseWriter: w}
	if errSize == 0 {
		res.SendOK(ComplaintController{Id: 1, Desc:"Hello V1"})
	} else {
		res.SendUnprocessableEntity(err) 
	} */
	
}