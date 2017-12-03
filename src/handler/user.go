package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"config"
	"datautils"
	"hashutils"
	"headerutils"
	"net/http"
	"user"

	"cloud.google.com/go/datastore"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func User(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodPost:
		userPost(w, r)

	default:
		http.NotFound(w, r)

	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func userPost(w http.ResponseWriter, r *http.Request) {

	u := new(user.User)
	err := datautils.JsonRequestBodyDecode(r, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, exist, err := emailExist(u.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if exist {
		http.Error(w, "Email has been used", http.StatusConflict)
		return
	}

	u.Password, err = hashutils.BcryptB64FromString(u.Password, hashutils.BcryptCostDefault)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := client.DatastoreClient.NewTransaction(client.Context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key := datastore.IncompleteKey(config.KindUser, nil)
	key.Namespace = config.NamespaceUser

	_, err = tx.Put(key, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	_, err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	headerutils.SetCookie(w, headerutils.CookieName, u.UID, headerutils.CookiePathRoot)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("/action"))

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func userDelete(w http.ResponseWriter, r *http.Request) {

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

	tx, err := client.DatastoreClient.NewTransaction(client.Context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tx.Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("/"))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
