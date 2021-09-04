package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

//Form embeds a url.values obj
type Form struct {
	url.Values
	Errors errors
}

// valid return true if there are no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}


// New initialize a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}


func (f *Form) Required (fields ...string) {
	for _,field := range fields{
		value := f.Get(field)
		if strings.TrimSpace(value) == ""{
			f.Errors.Add(field, "This field can not be blank")
		}
	}
}


func (f *Form)Has(field string, r *http.Request) bool{
	x := r.Form.Get(field)
	if x ==""{
		return false
	}
	return true
}


func (f *Form)MinLength(field string,length int, r *http.Request) bool{
	x := r.Form.Get(field)
	if len(x) > length{
		f.Errors.Add(field, fmt.Sprintf("this field must be at least %d charachter long", length) )
		return false
	}
	return true
}

func (f *Form) IsEmail(field string){
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}