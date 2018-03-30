package hashutils

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	BcryptCostMin     = bcrypt.MinCost
	BcryptCostMax     = bcrypt.MaxCost
	BcryptCostDefault = 12
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func BcryptB64FromString(password string, cost int) (string, error) {
	crypt, err := BcryptB64FromBytes([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return crypt, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func BcryptB64FromBytes(password []byte, cost int) (string, error) {
	crypt, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return "", err
	}

	cryptB64 := base64.StdEncoding.EncodeToString(crypt)

	return cryptB64, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func BcryptFromBytes(password []byte, cost int) ([]byte, error) {
	crypt, err := bcrypt.GenerateFromPassword(password, cost)

	if err != nil {
		return nil, err
	}

	return crypt, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func BcryptCompareB64Hash(b64PasswordHash string, password string) (bool, error) {
	passwordHash, err := base64.StdEncoding.DecodeString(b64PasswordHash)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
