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
	"github.com/mongodb/mongo-go-driver/mongo"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func User(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		userGet(w, r)

	case http.MethodPost:
		userPost(w, r)

	case http.MethodPut:
		userPut(w, r)

	default:
		http.NotFound(w, r)

	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func userGet(w http.ResponseWriter, r *http.Request) {

	uid, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	//keys, exist, err := uidExist(uid)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//if !exist {
	//http.Error(w, "UID does not exist", http.StatusBadRequest)
	//return
	//}

	//if len(keys) > 1 {
	//http.Error(w, "Multiple Keys for unique uid", http.StatusInternalServerError)
	//errorlog.ArchitectureLogicalError("Multiple keys for unique uid")
	//return
	//}

	//key := keys[0]

	ud := new(user.User)
	//err = client.DatastoreClient.Get(client.Context, key, ud)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	collection := client.MongoClient.Database(config.NamespaceAdmin).Collection(config.CollectionUser)

	filter := bson.NewDocument(bson.EC.Interface("uid", uid))
	err = collection.FindOne(context.Background(), filter).Decode(ud)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "UID does not exist", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeTemplate(w, "User", ud, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func userPost(w http.ResponseWriter, r *http.Request) {

	u := new(user.User)
	err := datautils.JsonRequestBodyDecode(r, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//_, exist, err := emailExist(u.Email)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//if exist {
	//http.Error(w, "Email has been used", http.StatusConflict)
	//return
	//}

	collection := client.MongoClient.Database(config.NamespaceAdmin).Collection(config.CollectionUser)

	ud := new(user.User)
	filter := bson.NewDocument(bson.EC.Interface("email", u.Email))
	err = collection.FindOne(context.Background(), filter).Decode(ud)
	if err != mongo.ErrNoDocuments {
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Error(w, "Email has been used", http.StatusConflict)
		return
	}

	u.Password, err = hashutils.BcryptB64FromString(u.Password, hashutils.BcryptCostDefault)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = collection.InsertOne(context.Background(), u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//tx, err := client.DatastoreClient.NewTransaction(client.Context)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//key := datastore.IncompleteKey(config.KindUser, nil)
	//key.Namespace = config.NamespaceUser

	//_, err = tx.Put(key, u)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusPreconditionFailed)
	//return
	//}

	//_, err = tx.Commit()
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusPreconditionFailed)
	//return
	//}

	headerutils.SetCookie(w, headerutils.CookieName, u.UID, headerutils.CookiePathRoot)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("/action"))

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func userPut(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	u := new(user.User)
	err = datautils.JsonRequestBodyDecode(r, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//keys, exist, err := emailExist(u.Email)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//if !exist {
	//http.Error(w, "Email does not exist", http.StatusBadRequest)
	//return
	//}

	collection := client.MongoClient.Database(config.NamespaceAdmin).Collection(config.CollectionUser)

	ud := new(user.User)
	filter := bson.NewDocument(bson.EC.Interface("email", u.Email))
	err = collection.FindOne(context.Background(), filter).Decode(ud)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Email does not exist", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	u.Password, err = hashutils.BcryptB64FromString(u.Password, hashutils.BcryptCostDefault)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//key := keys[0]

	//ud = new(user.User)
	//err = client.DatastoreClient.Get(client.Context, key, ud)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	if ud.UID != cookie {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	ud.Password = u.Password

	err = collection.FindOneAndReplace(context.Background(), filter, ud).Decode(&map[string]interface{}{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//tx, err := client.DatastoreClient.NewTransaction(client.Context)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//_, err = tx.Put(key, ud)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusPreconditionFailed)
	//return
	//}

	//_, err = tx.Commit()
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusPreconditionFailed)
	//return
	//}

	headerutils.SetCookie(w, headerutils.CookieName, ud.UID, headerutils.CookiePathRoot)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("/action"))

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func userDelete(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	u := new(user.User)
	err = datautils.JsonRequestBodyDecode(r, u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.MongoClient.Database(config.NamespaceAdmin).Collection(config.CollectionUser)

	ud := new(user.User)
	filter := bson.NewDocument(bson.EC.Interface("email", u.Email))
	err = collection.FindOne(context.Background(), filter).Decode(ud)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Email does not exist", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//keys, exist, err := emailExist(u.Email)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//if !exist {
	//http.Error(w, "Email does not exist", http.StatusBadRequest)
	//return
	//}

	//if len(keys) > 1 {
	//http.Error(w, "Multiple Keys for unique email", http.StatusInternalServerError)
	//return
	//}

	//key := keys[0]

	//ud := new(user.User)
	//err = client.DatastoreClient.Get(client.Context, key, ud)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	if ud.UID != cookie {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	err = collection.FindOneAndDelete(context.Background(), filter).Decode(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//tx, err := client.DatastoreClient.NewTransaction(client.Context)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = tx.Delete(key)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//_, err = tx.Commit()
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusPreconditionFailed)
	//return
	//}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("/"))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
