package encryptioner

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Encryptioner struct {}

func (e *Encryptioner) CreateHash(p string, i int) (string, error) {
	if i < 4 || 31 < i {
		return "", fmt.Errorf("%d is invalid value, expect number from 4 to 31", i)
	}

	password := []byte(p)
	hash, err := bcrypt.GenerateFromPassword(password, i)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
