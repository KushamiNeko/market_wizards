package headerutils

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

type EndPointUserAuth struct {
	Issuer string `json:"issuer"`
	ID     string `json:"id"`
	Email  string `json:"email"`
}

//////////////////////////////////////////////////////////////////////////////////////////////////////

func UserInfoFromEndPointHeaders(r *http.Request) (*EndPointUserAuth, error) {
	userInfoB64 := r.Header.Get("X-Endpoint-API-UserInfo")
	if userInfoB64 == "" {
		return nil, fmt.Errorf("no User Info in endpoint headers\n")
	}

	userInfoJson, err := base64.StdEncoding.DecodeString(userInfoB64)
	if err != nil {
		return nil, err
	}

	userInfo := new(EndPointUserAuth)

	err = json.Unmarshal(userInfoJson, userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
