package validator

import "regexp"

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Validator holds validation errors for input validation.
type Validator struct {
	Errors map[string]string
}

// New creates and returns a new Validator instance.
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// Valid returns true if there are no validation errors.
func (v *Validator) Valid() bool {
	return len(v.Errors) == 0
}

// AddError adds a new error message for a given key if it doesn't already exist.
func (v *Validator) AddError(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check runs a validation check and adds an error if the condition is false.
func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

// Matches checks if a string matches a given regular expression.
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
