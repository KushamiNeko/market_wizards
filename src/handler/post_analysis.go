package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"config"
	"context"
	"datautils"
	"fmt"
	"headerutils"
	"net/http"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func PostAnalysis(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		postAnalysisGet(w, r)

	case http.MethodPost:
		postAnalysisPost(w, r)

	case http.MethodPut:
		postAnalysisPut(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func postAnalysisGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "PostAnalysis", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func postAnalysisPost(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	objectRequest := new(datautils.PeriodSymbolStorage)

	err = datautils.JsonRequestBodyDecode(r, objectRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	paImg, err := datautils.FileReaderExtractImage(objectRequest.Data)
	if err != nil {
		http.Error(w, fmt.Sprintf("IBD Checkup: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	paDatastore := datautils.PeriodSymbolStorageNewBytes(objectRequest.From, objectRequest.To, objectRequest.Symbol, paImg.Bytes())

	collection := client.MongoClient.Database(config.NamespacePostAnalysis).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("from", paDatastore.From),
		bson.EC.Interface("to", paDatastore.To),
		bson.EC.Interface("symbol", paDatastore.Symbol),
	)

	err = collection.FindOne(context.Background(), filter).Decode(&datautils.DateSymbolStorage{})
	if err == nil {
		_, err = collection.ReplaceOne(context.Background(), filter, paDatastore)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(context.Background(), paDatastore)
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

func postAnalysisPut(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	objectRequest := new(datautils.PeriodSymbolStorage)

	err = datautils.JsonRequestBodyDecode(r, objectRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	paImg, err := datautils.FileReaderExtractImage(objectRequest.Data)
	if err != nil {
		http.Error(w, fmt.Sprintf("IBD Checkup: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	paDatastore := datautils.PeriodSymbolStorageNewBytes(objectRequest.From, objectRequest.To, objectRequest.Symbol, paImg.Bytes())

	collection := client.MongoClient.Database(config.NamespacePostAnalysis).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("from", paDatastore.From),
		bson.EC.Interface("to", paDatastore.To),
		bson.EC.Interface("symbol", paDatastore.Symbol),
	)

	err = collection.FindOne(context.Background(), filter).Decode(&datautils.DateSymbolStorage{})
	if err == nil {
		_, err = collection.ReplaceOne(context.Background(), filter, paDatastore)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(context.Background(), paDatastore)
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
