package common

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var (
	Validate *validator.Validate
)

func Init() {
	validate := validator.New()
	Validate = validate
}

func DoValidation(data interface{}) error {
	if err := Validate.Struct(data); err != nil {
		errs := err.(validator.ValidationErrors)
		messages := make([]string, 0)
		for _, e := range errs {
			messages = append(messages, fmt.Sprintf("%v %v", e.Field(), e.Tag()))
		}
		newErr := errors.New(strings.Join(messages, ", "))
		return newErr
	}

	return nil
}
