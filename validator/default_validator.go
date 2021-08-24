package validator

import (
	"regexp"
)

type DefaultPasswordValidator struct {
	Password string
}

func NewDefaultPasswordValidator() Validator{
	return &DefaultPasswordValidator{}
}

func (v *DefaultPasswordValidator) Validate(p string) bool {
	lowercase := regexp.MustCompile(`.*?[a-z]+`)
	uppercase := regexp.MustCompile(`.*?[A-Z]+`)
	number    := regexp.MustCompile(`.*?\d+`)
	mix       := regexp.MustCompile(`^([a-z]|\d|[A-Z]){8,100}$`)

	if result := lowercase.MatchString(p); !result {
		return false
	}

	if result := uppercase.MatchString(p); !result {
		return false
	}

	if result := number.MatchString(p); !result {
		return false
	}

	if result := mix.MatchString(p); !result {
		return false
	}

	return true
}
