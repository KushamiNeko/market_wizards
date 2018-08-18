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
	"ibd"
	"net/http"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func IBD(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	//case http.MethodGet:
	//ibdGet(w, r)

	case http.MethodPost:
		ibdPost(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ibdGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func ibdPost(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	objectRequest := new(datautils.DateSymbolStorage)

	err = datautils.JsonRequestBodyDecode(r, objectRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ibdBuffer, err := datautils.FileReaderExtract(objectRequest.Data)
	if err != nil {
		http.Error(w, fmt.Sprintf("IBD Checkup: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	ibdCheckup, err := ibd.Parse(ibdBuffer)
	if err != nil {
		http.Error(w, fmt.Sprintf("IBD Checkup: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	ibdJson, err := json.Marshal(ibdCheckup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ibdDatastore := datautils.DateSymbolStorageNewBytes(objectRequest.Date, objectRequest.Symbol, ibdJson)

	collection := client.MongoClient.Database(config.NamespaceIBD).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("date", ibdDatastore.Date),
		bson.EC.Interface("symbol", ibdDatastore.Symbol),
	)

	err = collection.FindOne(context.Background(), filter).Decode(&datautils.DateSymbolStorage{})
	if err == nil {
		_, err = collection.ReplaceOne(context.Background(), filter, ibdDatastore)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(context.Background(), ibdDatastore)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
