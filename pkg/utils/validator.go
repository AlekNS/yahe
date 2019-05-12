package utils

import validator "gopkg.in/go-playground/validator.v9"

// MustValidate helps to validate struct.
func MustValidate(val interface{}) {
	var err = validator.New().Struct(val)

	if err != nil {
		panic(err)
	}
}
