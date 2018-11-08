package controllers

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

	res.SendOK(ComplaintController{Id: 3, Desc:"Hello V3"})
}