package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data, errors(map[string][]string{}),
	}
}

func (f *Form) Has(field string, r *http.Request) bool {
	formField := f.Get(field)
	if formField == "" {
		return false
	}
	return true
}
