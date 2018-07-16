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
	"sort"
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

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
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

	collection := client.MongoClient.Database(config.NamespaceWatchList).Collection(cookie)

	var filter *bson.Document

	if symbol == "" {
		filter = bson.NewDocument()
	} else {
		filter = bson.NewDocument(
			bson.EC.Interface("symbol", symbol),
		)
	}

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	items := make([]*watchlist.WatchListItem, 0)

	for cursor.Next(context.Background()) {

		t := new(watchlist.WatchListItem)

		err := cursor.Decode(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		items = append(items, t)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sort.Slice(items, func(i, j int) bool {
		ii := items[i]
		ij := items[j]

		iis := fmt.Sprintf("%v%v%v%v", ii.Priority, ii.Status, ii.Fundamentals, ii.Symbol)
		ijs := fmt.Sprintf("%v%v%v%v", ij.Priority, ij.Status, ij.Fundamentals, ij.Symbol)

		return iis < ijs
	})

	wt.Items = items

	if capital == "" || size == "" {
		writeTemplate(w, "WatchList", wt, nil)
		return
	}

	capitalF, err := strconv.ParseFloat(capital, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sizeF, err := strconv.ParseFloat(size, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	wt.PositionSize = fmt.Sprintf("%.2f%%", sizeF)

	for _, t := range items {
		t.Caculate(capitalF, sizeF)
	}

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

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	symbol := r.URL.Query().Get("Symbol")

	if symbol == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.MongoClient.Database(config.NamespaceWatchList).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("symbol", symbol),
	)

	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
