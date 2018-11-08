package v2

import (
	"net/http"

	"github.com/janaonline/icmyc/app/handlers"
)

type ComplaintController struct{
	Id int `json:"id,omitempty"`
	Desc string `json:"body,omitempty"`

}

func (c *ComplaintController) Index(w http.ResponseWriter, r *http.Request)  {
	res := handlers.Response{ResponseWriter: w}

	res.SendOK(ComplaintController{Id: 2, Desc:"Hello V2"})
}