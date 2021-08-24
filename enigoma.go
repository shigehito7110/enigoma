package enigoma

import(
	"fmt"
	vl "github.com/shigehito7110/enigoma/validator" // validator package for password
	en "github.com/shigehito7110/enigoma/encryptioner" // package to create hash of password
)

type Enigoma struct {
	Password 		 string
	Validator 	 vl.Validator
	Encryptioner *en.Encryptioner
}

type validator func() vl.Validator

func NewEnigoma(p string, vs ...validator) (*Enigoma, error){
	en := new(en.Encryptioner)
	e := &Enigoma{
		Password: p,
		Encryptioner: en,
	}

	switch(len(vs)) {
		case 0:
			e.Validator = vl.NewDefaultPasswordValidator()
		case 1:
			for _, v := range vs {
				e.Validator = v()
			}
		default:
			return nil, fmt.Errorf("too many arguments, given %d expected 1", len(vs))
	}

	return e, nil
}

func (e *Enigoma) Validate() error {
	if result := e.Validator.Validate(e.Password); !result {
		return fmt.Errorf("%s is invalid values for password", e.Password)
	}

	return nil
}

func (e *Enigoma) CreateHash(_i ...int) (string, error) {
	var i int
	switch(len(_i)) {
	case 0:
		i = 10
	case 1:
		for _, val := range _i {
			i = val
		}
	default:
		return "", fmt.Errorf("too many arguments, given %d expected 1", len(_i))
	}

	h, err := e.Encryptioner.CreateHash(e.Password, i)
	if err != nil {
		//errorをwrapする
		return "", fmt.Errorf("%s coudn't create hash", e.Password)
	}

	return h, nil
}

func (e *Enigoma) Run(_i ...int) (string, error) {
	if err := e.Validate(); err != nil {
		return "", err
	}

	h, err := e.CreateHash(_i...)
	if err != nil {
		return "", err
	}

	return h, nil
}
