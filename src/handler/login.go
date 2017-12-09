package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"datautils"
	"hashutils"
	"headerutils"
	"net/http"
	"user"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Login(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		loginGet(w, r)

	case http.MethodPost:
		loginPost(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func loginGet(w http.ResponseWriter, r *http.Request) {

	writeTemplate(w, "Login", nil, func() {
		headerutils.DeleteCookie(w, headerutils.CookieName, headerutils.CookiePathRoot)
	})
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func loginPost(w http.ResponseWriter, r *http.Request) {

	u := new(user.User)
	err := datautils.JsonRequestBodyDecode(r, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	keys, exist, err := emailExist(u.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !exist {
		http.Error(w, "Email does not exist", http.StatusBadRequest)
		return
	}

	if len(keys) > 1 {
		http.Error(w, "Multiple Keys for unique email", http.StatusInternalServerError)
		return
	}

	key := keys[0]

	ud := new(user.User)
	err = client.DatastoreClient.Get(client.Context, key, ud)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	match, err := hashutils.BcryptCompareWithB64PasswordHash(ud.Password, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !match {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	headerutils.SetCookie(w, headerutils.CookieName, ud.UID, headerutils.CookiePathRoot)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("/action"))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
