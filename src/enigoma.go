package enigoma

import(
	"fmt"
	vl "github.com/shigehito7110/enigoma/validator"
)

type Enigoma struct {
	Password string
	Validator vl.Validator
}

type validator func() vl.Validator

func NewEnigoma(p string, vs ...validator) (*Enigoma, error){
	e := &Enigoma{
		Password: p,
	}

	switch(len(vs)) {
		case 0:
			e.Validator = vl.NewDefaultPasswordValidator()
		case 1:
			for _, v := range vs {
				e.Validator = v()
			}
		default:
			return nil, fmt.Errorf("Too many arguments,given %d expected 1", len(vs))
	}

	return e, nil
}

func (e *Enigoma) Validate() error {
	if result := e.Validator.Validate(e.Password); result {
		return nil
	}

	return fmt.Errorf("%s is invalid values for password", e.Password)
}
