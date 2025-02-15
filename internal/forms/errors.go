package forms

type errors map[string][]string

func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

func (e errors) Get(field string) string {
	errorSting := e[field]
	if len(errorSting) == 0 {
		return ""
	}
	return errorSting[0]
}
