package encryptioner

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Encryptioner struct {}

func (e *Encryptioner) CreateHash(p string, i int) (string, error) {
	if i < 4 || 31 < i {
		return "", fmt.Errorf("%d is invalid value, expect number include in 4-31", i)
	}

	password  := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, i)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (e *Encryptioner) Check(hp, p string) error {
	password     := []byte(p)
	hashpassword := []byte(hp)
	if err := bcrypt.CompareHashAndPassword(hashpassword, password); err != nil {
		return fmt.Errorf("%s is not correct password", p)
	}

	return nil
}
