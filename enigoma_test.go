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

func TestEnigomaCreateHash(t *testing.T) {
	e, _ := NewEnigoma(correctPassForDef)
	t.Run("Encryptioner/", func(t *testing.T) {
		t.Run("DefaultValue", func(t *testing.T) {
			if _, err := e.CreateHash(); err != nil {
				t.Errorf("Return error")
			}
		})

		t.Run("CustomValue", func(t *testing.T) {
			t.Run("Success", func(t *testing.T) {
				// TODO(shige): なぜかループ処理が走らないので修正後追加
				// for i := 4; i <= 31; i++ {
				// 	if _, err := e.CreateHash(i); err != nil {
				// 		t.Errorf("Return error")
				// 	}
				// }
				if _, err := e.CreateHash(5); err != nil {
					t.Errorf("Return error")
				}
			})

			t.Run("Failure", func(t *testing.T) {
				t.Run("numeber not include in 4-31", func(t *testing.T) {
					if _, err := e.CreateHash(1); err == nil {
						t.Errorf("Not return error")
					}

					if _, err := e.CreateHash(32); err == nil {
						t.Errorf("Not return error")
					}

					if _, err := e.CreateHash(1); err.Error() != "1 is invalid value, expect number include in 4-31" {
						t.Errorf("Not correct error message")
					}

					if _, err := e.CreateHash(32); err.Error() != "32 is invalid value, expect number include in 4-31" {
						t.Errorf("Not correct error message")
					}
				})

				t.Run("too  many arguments", func(t *testing.T) {
					if _ ,err := e.CreateHash(3,4); err == nil {
						t.Errorf("Not return error")
					}
				})
			})
		})
	})
}
