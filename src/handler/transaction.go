package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"config"
	"context"
	"datautils"
	"encoding/json"
	"headerutils"
	"net/http"
	"strconv"
	"transaction"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Transaction(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		dateOfPurchase := r.URL.Query().Get("DateOfPurchase")
		symbol := r.URL.Query().Get("Symbol")

		if dateOfPurchase == "" || symbol == "" {
			transactionGet(w, r)
		} else {
			transactionSearch(w, r)
		}

	case http.MethodPost:
		transactionPost(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func transactionGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "TransactionNew", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func transactionSearch(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	dateOfPurchase := r.URL.Query().Get("DateOfPurchase")
	symbol := r.URL.Query().Get("Symbol")

	dateOfPurchaseI, err := strconv.ParseInt(dateOfPurchase, 10, 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.MongoClient.Database(config.NamespaceTransaction).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("order", "buy"),
		bson.EC.Interface("date", int(dateOfPurchaseI)),
		bson.EC.Interface("symbol", symbol),
	)

	t := new(transaction.BuyOrder)

	err = collection.FindOne(context.Background(), filter).Decode(t)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "No Document Found", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	headerutils.ContentTypeJsonUTF8(w)

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func transactionPost(w http.ResponseWriter, r *http.Request) {

	order := r.URL.Query().Get("Order")
	if order == "" {
		http.Error(w, "Invalid Order Type", http.StatusBadRequest)
		return
	}

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	var t transaction.Order

	if order == "buy" {
		t = new(transaction.BuyOrder)
	}

	if order == "sell" {
		t = new(transaction.SellOrder)
	}

	err = datautils.JsonRequestBodyDecode(r, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//if t.GetJsonIBDCheckup() == "" {
	//http.Error(w, "Missing IBD Checkup File", http.StatusBadRequest)
	//return
	//}

	//ibdBuffer, err := datautils.FileReaderExtract(t.GetJsonIBDCheckup())
	//if err != nil {
	//http.Error(w, fmt.Sprintf("IBD Checkup: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//ibdCheckup, err := ibd.Parse(ibdBuffer)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("IBD Checkup: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//ibdJson, err := json.Marshal(ibdCheckup)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//ibdDatastore := ibd.IBDCheckupDatastoreNew(t.GetDate(), t.GetSymbol(), ibdJson)
	//ibdDatastore := datautils.DataIDStorageNewBytes(t.GetIBDCheckupID(), ibdJson)

	collection := client.MongoClient.Database(config.NamespaceTransaction).Collection(cookie)

	_, err = collection.InsertOne(context.Background(), t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//collection = client.MongoClient.Database(config.NamespaceIBD).Collection(cookie)

	//_, err = collection.InsertOne(context.Background(), ibdDatastore)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("/action"))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func transactionPut(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func transactionDelete(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
