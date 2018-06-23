package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"config"
	"context"
	"datautils"
	"encoding/json"
	"fmt"
	"headerutils"
	"marketsmith"
	"net/http"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func MarketSmith(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	//case http.MethodGet:
	//marketsmithGet(w, r)

	case http.MethodPost:
		marketsmithPost(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func marketsmithGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	//writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func marketsmithPost(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	objectRequest := new(ObjectRequestBody)

	err = datautils.JsonRequestBodyDecode(r, objectRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	msBuffer, err := datautils.FileReaderExtract(objectRequest.Object)
	if err != nil {
		http.Error(w, fmt.Sprintf("MarketSmith: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	ms, err := marketsmith.Parse(msBuffer)
	if err != nil {
		http.Error(w, fmt.Sprintf("MarketSmith: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	msJson, err := json.Marshal(ms)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msDatastore := datautils.DataIDStorageNewBytes(objectRequest.DateSymbolID(), msJson)

	collection := client.MongoClient.Database(config.NamespaceMarketSmith).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("id", msDatastore.ID),
	)

	err = collection.FindOne(context.Background(), filter).Decode(&datautils.DataIDStorage{})
	if err == nil {

	} else {
		if err == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(context.Background(), msDatastore)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	//_, err = collection.InsertOne(context.Background(), marketsmithDatastore)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
