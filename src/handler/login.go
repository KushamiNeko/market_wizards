package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"config"
	"context"
	"datautils"
	"hashutils"
	"headerutils"
	"net/http"
	"user"

	"github.com/mongodb/mongo-go-driver/bson"
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

	ud := new(user.User)

	collection := client.MongoClient.Database(config.NamespaceAdmin).Collection(config.CollectionUser)

	filter := bson.NewDocument(bson.EC.Interface("email", u.Email))
	err = collection.FindOne(context.Background(), filter).Decode(ud)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	match, err := hashutils.BcryptCompareB64Hash(ud.Password, u.Password)
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
