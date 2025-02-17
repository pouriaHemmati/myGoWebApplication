package forms

import (
	"github.com/asaskevich/govalidator"
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

func (f *Form) MinLength(fields string, length int, r *http.Request) bool {
	actualLength := f.Get(fields)
	if len(strings.TrimSpace(actualLength)) < length {
		f.Errors.Add(fields, "This field must be at least %d characters.")
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "This field must be a valid email address.")
	}
}
