package user

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"encoding/json"
	"hashutils"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

type User struct {
	Email string

	Password string

	UID string
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func (u *User) JsonDecode(buffer []byte) error {

	err := json.Unmarshal(buffer, u)
	if err != nil {
		return err
	}

	u.UID = hashutils.RandBytesGenerateB64(config.KeyLengthDefault)

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
