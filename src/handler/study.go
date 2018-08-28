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

func ChartsStudy(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		chartsStudyGet(w, r)

	case http.MethodPost:
		chartsStudyPost(w, r)

	case http.MethodPut:
		chartsStudyPut(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func chartsStudyGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	//writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func chartsStudyPost(w http.ResponseWriter, r *http.Request) {

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

	chartImg, err := datautils.FileReaderExtractImage(objectRequest.Data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Charts Study: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	chartDatastore := datautils.DateSymbolStorageNewBytes(objectRequest.Date, objectRequest.Symbol, chartImg.Bytes())

	collection := client.MongoClient.Database(config.NamespaceChartsStudy).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("date", chartDatastore.Date),
		bson.EC.Interface("symbol", chartDatastore.Symbol),
	)

	err = collection.FindOne(context.Background(), filter).Decode(&datautils.DateSymbolStorage{})
	if err == nil {
		_, err = collection.ReplaceOne(context.Background(), filter, chartDatastore)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(context.Background(), chartDatastore)
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

func chartsStudyPut(w http.ResponseWriter, r *http.Request) {

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

	chartImg, err := datautils.FileReaderExtractImage(objectRequest.Data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Charts Study: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	chartDatastore := datautils.DateSymbolStorageNewBytes(objectRequest.Date, objectRequest.Symbol, chartImg.Bytes())

	collection := client.MongoClient.Database(config.NamespaceChartsStudy).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("date", chartDatastore.Date),
		bson.EC.Interface("symbol", chartDatastore.Symbol),
	)

	err = collection.FindOne(context.Background(), filter).Decode(&datautils.DateSymbolStorage{})
	if err == nil {
		_, err = collection.ReplaceOne(context.Background(), filter, chartDatastore)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(context.Background(), chartDatastore)
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
