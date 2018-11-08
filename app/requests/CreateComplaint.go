package requests

import (
	"net/http"
	"net/url"
	"github.com/thedevsaddam/govalidator"
	"github.com/go-ozzo/ozzo-validation"
	_ "github.com/go-ozzo/ozzo-validation/is"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

func (a Address) Validate() error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Street, validation.Required, validation.Length(5, 50)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&a.City, validation.Required, validation.Length(5, 50)),
		// State cannot be empty, and must be a string consisting of two letters in upper case
		validation.Field(&a.State, validation.Required),
		// State cannot be empty, and must be a string consisting of five digits
		validation.Field(&a.Zip, validation.Required),
	)
}

func CreateComplaintRequest(r *http.Request) url.Values{
	rules := govalidator.MapData{
		"username": []string{"required", "between:3,8"},
	}

	messages := govalidator.MapData{
		"username": []string{"required:আপনাকে অবশ্যই ইউজারনেম দিতে হবে", "between:ইউজারনেম অবশ্যই ৩-৮ অক্ষর হতে হবে"},
		"phone":    []string{"digits:ফোন নাম্বার অবশ্যই ১১ নম্বারের হতে হবে"},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	v := govalidator.New(opts)

	return v.Validate()
}


