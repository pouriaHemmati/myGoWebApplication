package forms

import (
	"net/http"
	"net/url"
	"strings"
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

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Has(field string, r *http.Request) bool {
	formField := f.Get(field)
	if formField == "" {
		f.Errors.Add(field, "This field cannot be empty.")

		return false
	}
	return true
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if len(strings.TrimSpace(value)) == 0 {
			f.Errors.Add(field, "This field cannot be empty.")
		}
	}
}
