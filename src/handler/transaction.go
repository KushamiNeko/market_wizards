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

	//var q *datastore.Query
	//var entities []datastore.PropertyList

	//q = datastore.NewQuery(cookie).Namespace(config.NamespaceTransaction)
	//q = q.Filter("Order =", "buy")
	//q = q.KeysOnly()

	//orderKeys, err := client.DatastoreClient.GetAll(client.Context, q, &entities)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//if len(orderKeys) <= 0 {
	//http.Error(w, "No orders keys", http.StatusBadRequest)
	//return
	//}

	//d, err := strconv.ParseInt(dateOfPurchase, 10, 32)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusBadRequest)
	//return
	//}

	//q = datastore.NewQuery(cookie).Namespace(config.NamespaceTransaction)
	//q = q.Filter("Date =", int(d))
	//q = q.KeysOnly()

	//dateKeys, err := client.DatastoreClient.GetAll(client.Context, q, &entities)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//if len(dateKeys) <= 0 {
	//http.Error(w, "No date keys", http.StatusBadRequest)
	//return
	//}

	//q = datastore.NewQuery(cookie).Namespace(config.NamespaceTransaction)
	//q = q.Filter("Symbol =", symbol)
	//q = q.KeysOnly()

	//symbolKeys, err := client.DatastoreClient.GetAll(client.Context, q, &entities)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//if len(symbol) <= 0 {
	//http.Error(w, "No symbol keys", http.StatusBadRequest)
	//return
	//}

	//var key *datastore.Key

	//for _, ok := range orderKeys {

	//matchDate := false
	//matchSymbol := false

	//for _, dk := range dateKeys {
	//if ok.Equal(dk) {
	//matchDate = true
	//break
	//}
	//}

	//for _, sk := range symbolKeys {
	//if ok.Equal(sk) {
	//matchSymbol = true
	//break
	//}
	//}

	//if matchDate && matchSymbol {
	//key = ok
	//break
	//}
	//}

	//t := new(transaction.Order)

	//err = client.DatastoreClient.Get(client.Context, key, t)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

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

	//t := new(transaction.Order)
	err = datautils.JsonRequestBodyDecode(r, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if t.GetJsonIBDCheckup() == "" {
		http.Error(w, "Missing IBD Checkup File", http.StatusBadRequest)
		return
	}

	ibdBuffer, err := datautils.FileReaderExtract(t.GetJsonIBDCheckup())
	if err != nil {
		http.Error(w, fmt.Sprintf("IBD Checkup: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	ibdCheckup, err := ibd.Parse(ibdBuffer)
	if err != nil {
		http.Error(w, fmt.Sprintf("IBD Checkup: %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	//if t.JsonChartD == "" {
	//http.Error(w, "Missing Daily Chart", http.StatusBadRequest)
	//return
	//}

	//chartDBuffer, err := datautils.FileReaderExtractImage(t.JsonChartD)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("Daily Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//if t.JsonChartW == "" {
	//http.Error(w, "Missing Weekly Chart", http.StatusBadRequest)
	//return
	//}

	//chartWBuffer, err := datautils.FileReaderExtractImage(t.JsonChartW)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("Weekly Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//if t.JsonChartNDQCD == "" {
	//http.Error(w, "Missing NDQC Daily Chart", http.StatusBadRequest)
	//return
	//}

	//chartNdqcDBuffer, err := datautils.FileReaderExtractImage(t.JsonChartNDQCD)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("NDQC Daily Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//if t.JsonChartNDQCW == "" {
	//http.Error(w, "Missing NDQC Weekly Chart", http.StatusBadRequest)
	//return
	//}

	//chartNdqcWBuffer, err := datautils.FileReaderExtractImage(t.JsonChartNDQCW)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("NDQC Weekly Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//if t.JsonChartSP5D == "" {
	//http.Error(w, "Missing S&P5 Daily Chart", http.StatusBadRequest)
	//return
	//}

	//chartSp5DBuffer, err := datautils.FileReaderExtractImage(t.JsonChartSP5D)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("S&P5 Daily Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//if t.JsonChartSP5W == "" {
	//http.Error(w, "Missing S&P5 Weekly Chart", http.StatusBadRequest)
	//return
	//}

	//chartSp5WBuffer, err := datautils.FileReaderExtractImage(t.JsonChartSP5W)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("S&P5 Weekly Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//if t.JsonChartNYCD == "" {
	//http.Error(w, "Missing NYC Daily Chart", http.StatusBadRequest)
	//return
	//}

	//chartNycDBuffer, err := datautils.FileReaderExtractImage(t.JsonChartNYCD)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("NYC Daily Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//if t.JsonChartNYCW == "" {
	//http.Error(w, "Missing NYC Weekly Chart", http.StatusBadRequest)
	//return
	//}

	//chartNycWBuffer, err := datautils.FileReaderExtractImage(t.JsonChartNYCW)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("NYC Weekly Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//if t.JsonChartDJIAD == "" {
	//http.Error(w, "Missing DJIA Daily Chart", http.StatusBadRequest)
	//return
	//}

	//chartDjiaDBuffer, err := datautils.FileReaderExtractImage(t.JsonChartDJIAD)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("DJIA Daily Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//if t.JsonChartDJIAW == "" {
	//http.Error(w, "Missing DJIA Weekly Chart", http.StatusBadRequest)
	//return
	//}

	//chartDjiaWBuffer, err := datautils.FileReaderExtractImage(t.JsonChartDJIAW)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("DJIA Weekly Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//if t.JsonChartRUSD == "" {
	//http.Error(w, "Missing RUS Daily Chart", http.StatusBadRequest)
	//return
	//}

	//chartRusDBuffer, err := datautils.FileReaderExtractImage(t.JsonChartRUSD)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("RUS Daily Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//if t.JsonChartRUSW == "" {
	//http.Error(w, "Missing RUS Weekly Chart", http.StatusBadRequest)
	//return
	//}

	//chartRusWBuffer, err := datautils.FileReaderExtractImage(t.JsonChartRUSW)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("RUS Weekly Chart: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	ibdJson, err := json.Marshal(ibdCheckup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ibdDatastore := ibd.IBDCheckupDatastoreNew(t.GetDate(), t.GetSymbol(), ibdJson)

	collection := client.MongoClient.Database(config.NamespaceTransaction).Collection(cookie)

	_, err = collection.InsertOne(context.Background(), t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	collection = client.MongoClient.Database(config.NamespaceIBD).Collection(cookie)

	_, err = collection.InsertOne(context.Background(), ibdDatastore)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//tx, err := client.DatastoreClient.NewTransaction(client.Context)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//tkey := datastore.IncompleteKey(cookie, nil)
	//tkey.Namespace = config.NamespaceTransaction

	//iKey, _, err := ibdGetKey(cookie, ibdDatastore.ID)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//_, err = tx.Put(tkey, t)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//_, err = tx.Put(iKey, ibdDatastore)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//_, err = tx.Commit()
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusPreconditionFailed)
	//return
	//}

	//cookiePath := url.PathEscape(cookie)

	//err = writeStorageObject(
	//filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, t.Symbol)), chartDBuffer)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = writeStorageObject(filepath.Join(
	//cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, t.Symbol)), chartWBuffer)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = writeStorageObject(filepath.Join(
	//cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, "NDQC")), chartNdqcDBuffer)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = writeStorageObject(filepath.Join(
	//cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, "NDQC")), chartNdqcWBuffer)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = writeStorageObject(filepath.Join(
	//cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, "S&P5")), chartSp5DBuffer)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = writeStorageObject(filepath.Join(
	//cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, "S&P5")), chartSp5WBuffer)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = writeStorageObject(
	//filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, "NYC")), chartNycDBuffer)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = writeStorageObject(
	//filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, "NYC")), chartNycWBuffer)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = writeStorageObject(
	//filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, "DJIA")), chartDjiaDBuffer)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = writeStorageObject(
	//filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, "DJIA")), chartDjiaWBuffer)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = writeStorageObject(
	//filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, "RUS")), chartRusDBuffer)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//err = writeStorageObject(
	//filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, "RUS")), chartRusWBuffer)
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
