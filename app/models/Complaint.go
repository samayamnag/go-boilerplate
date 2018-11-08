package models

import (

)

type Complaint struct {
	ID  int `json:"id,omitempty"`
	Descritpion string `json:"body,omitemtpy"`
}