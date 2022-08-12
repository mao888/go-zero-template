package validator

import (
	"errors"
	"github.com/go-playground/validator"
	"github.com/zeromicro/go-zero/core/logx"
)

const (
	isValidAgeTag   = "isValidAge"
	isValidHobbyTag = "isValidHobby"
)

var (
	InvalidAgeErr   = errors.New("age必须在0到200")
	InvalidHobbyErr = errors.New("hobby should be 'ball' or 'swim'")
)

func validAge(fl validator.FieldLevel) bool {
	age := fl.Field().Int()
	if age >= 0 && age <= 200 {
		return true
	}
	return false
}

func validHobby(fl validator.FieldLevel) bool {
	hobby := fl.Field().String()
	if hobby == "ball" || hobby == "swim" {
		return true
	}
	return false
}

var validateObj *validator.Validate

func registerValidation(tag string, fn validator.Func) {
	if err := validateObj.RegisterValidation(tag, fn); err != nil {
		logx.Errorf("register validator for '%s' error: %v", tag, err)
	}
}

func init() {
	validateObj = validator.New()

	registerValidation(isValidAgeTag, validAge)
	registerValidation(isValidHobbyTag, validHobby)
}

func ValidateGreet(s interface{}) error {
	err := validateObj.Struct(s)
	if err == nil {
		return nil
	}

	// refer to: https://github.com/go-playground/validator/issues/559
	for _, validateErr := range err.(validator.ValidationErrors) {
		// only check first error
		switch validateErr.Tag() {
		case isValidAgeTag:
			// return custom error message
			return InvalidAgeErr
		case isValidHobbyTag:
			return InvalidHobbyErr
		default:
			return err
		}
	}

	return err
}
