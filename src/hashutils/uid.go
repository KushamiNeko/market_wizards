package hashutils

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func GenerateUID(user, password string) string {
	UID :=
		ShaB64FromString(user, password, config.UIDSecret)

	return UID
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
