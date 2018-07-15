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
	"strconv"
	"watchlist"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type watchListTemplate struct {
	Capital string
	Size    string
	Symbol  string

	PositionSize string

	Items []*watchlist.WatchListItem
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func WatchList(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		action := r.URL.Query().Get("Action")

		if action == "" {
			watchListGet(w, r)
		} else if action == "new" {
			watchListNewGet(w, r)
		} else {
			http.NotFound(w, r)
		}

	case http.MethodPost:
		watchListPost(w, r)

	case http.MethodDelete:
		watchListDelete(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func watchListGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	capital := r.URL.Query().Get("Capital")
	size := r.URL.Query().Get("Size")
	symbol := r.URL.Query().Get("Symbol")

	wt := new(watchListTemplate)
	wt.Capital = capital
	wt.Size = size

	if capital == "" || size == "" {
		writeTemplate(w, "WatchList", wt, nil)
		return
	}

	if symbol == "" {

	}

	//capitalF, err := strconv.ParseFloat(capital, 64)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusBadRequest)
	//return
	//}

	sizeF, err := strconv.ParseFloat(size, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wt.PositionSize = fmt.Sprintf("%.2f%%", sizeF)

	writeTemplate(w, "WatchList", wt, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func watchListNewGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "WatchListNew", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func watchListPost(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	objectRequest := new(watchlist.WatchListItem)

	err = datautils.JsonRequestBodyDecode(r, objectRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.MongoClient.Database(config.NamespaceWatchList).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("symbol", objectRequest.Symbol),
	)

	err = collection.FindOne(context.Background(), filter).Decode(&datautils.DateSymbolStorage{})
	if err == nil {
		_, err = collection.ReplaceOne(context.Background(), filter, objectRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		if err == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(context.Background(), objectRequest)
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

func watchListDelete(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	//writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
