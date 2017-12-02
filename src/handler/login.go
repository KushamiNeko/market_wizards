package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"client"
	"config"
	"datautils"
	"hashutils"
	"minify"
	"net/http"
	"user"

	"cloud.google.com/go/datastore"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Login(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		buffer := new(bytes.Buffer)

		err := templates.ExecuteTemplate(
			buffer,
			"Login",
			nil,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(minify.Minify(buffer.Bytes()))

	case http.MethodPost:
		return

	default:
		http.NotFound(w, r)
		return
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewUser(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		u := new(user.User)

		err := datautils.JsonRequestBodyDecode(r, u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		u.Password, err = hashutils.BcryptB64FromString(u.Password, hashutils.BcryptCostDefault)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tx := client.DatastoreClient.NewTransaction(client.Context)

		key := datastore.IncompleteKey(config.KindUser, nil)
		key.Namespace = config.NamespaceUser

		_, err := tx.Put(key, u)
		if err != nil {
			http.Error(w, err.Error(), http.StatusPreconditionFailed)
			return
		}

		_, err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		//buffer := new(bytes.Buffer)

		//err := templates.ExecuteTemplate(
		//buffer,
		//"Login",
		//nil,
		//)
		//if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return
		//}

		//w.Write(minify.Minify(buffer.Bytes()))

	default:
		http.NotFound(w, r)
		return
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
