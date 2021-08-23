package enigoma

import (
	"fmt"
	"testing"
	vl "github.com/shigehito7110/enigoma/validator"
)

// Success validator struct
type testV struct {}
func (t *testV) Validate(p string) bool {
	return true
}

const correctPassForDef = "Password1234"

func TestNewEnigoma(t *testing.T) {
	f1 := func() vl.Validator {
		return new(testV)
	}

	f2 := func() vl.Validator {
		return new(testV)
	}

	t.Run("DefaultValidatorFunc/Success", func(t *testing.T) {
		if _, err := NewEnigoma(correctPassForDef); err != nil {
			t.Errorf("Return error")
		}
	})

	t.Run("CustomValidatorFunc", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			if _, err := NewEnigoma(correctPassForDef, f1); err != nil {
				t.Errorf("Return error")
			}
		})

		t.Run("Failure", func(t *testing.T) {
			if _, err := NewEnigoma(correctPassForDef, f1, f2); err == nil {
				t.Errorf("No return error")
			}

			if _, err := NewEnigoma(correctPassForDef, f1, f2); err.Error() != "too many arguments, given 2 expected 1" {
				t.Errorf("Not correct error message")
			}
		})
	})
}

func TestEnigomaValidate(t *testing.T) {
	t.Run("DefaultValidatorFunc", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			e, _ := NewEnigoma(correctPassForDef)
			if err := e.Validate(); err != nil {
				t.Errorf("Return error")
			}
		})

		t.Run("Failure", func(t *testing.T) {
			// InvalidPasswords
			var passwords []string = []string{
				"password",
				"12345",
				"PASSWORD",
				"Pass1",
				"Password1234#",
			}

			for _, pass := range passwords {
				e, _ := NewEnigoma(pass)
				if err := e.Validate(); err == nil {
					t.Errorf("Not return error")
				}

				if err := e.Validate(); err.Error() != fmt.Sprintf("%s is invalid values for password", pass) {
					t.Errorf("Not correct error message")
				}
			}
		})
	})
}
