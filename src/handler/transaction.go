package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"client"
	"config"
	"datautils"
	"encoding/json"
	"fmt"
	"headerutils"
	"ibd"
	"net/http"
	"net/url"
	"path/filepath"
	"transaction"

	"cloud.google.com/go/datastore"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Transaction(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		transactionGet(w, r)

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

func transactionPost(w http.ResponseWriter, r *http.Request) {

	var buffer *bytes.Buffer

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	t := new(transaction.Order)
	err = datautils.JsonRequestBodyDecode(r, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//c.ID = t.IBDCheckup

	tx, err := client.DatastoreClient.NewTransaction(client.Context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tkey := datastore.IncompleteKey(cookie, nil)
	tkey.Namespace = config.NamespaceTransaction

	//ckey := datastore.IncompleteKey(cookie, nil)
	//ckey.Namespace = config.NamespaceIBD

	_, err = tx.Put(tkey, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//_, err = tx.Put(ckey, c)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	_, err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	if t.JsonIBDCheckup == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtract(t.JsonIBDCheckup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ibd, err := ibd.Parse(buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ibdJson, err := json.Marshal(ibd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	buffer = bytes.NewBuffer(ibdJson)

	cookiePath := url.PathEscape(cookie)

	err = writeStorageObject(
		filepath.Join(cookiePath, config.StorageNamespaceIBDs, fmt.Sprintf("%d_%s", t.Date, t.Symbol)), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartD == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(
		filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, t.Symbol)), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartW == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartW)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(filepath.Join(
		cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, t.Symbol)), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartNDQCD == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartNDQCD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(filepath.Join(
		cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, "NDQC")), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartNDQCW == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartNDQCW)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(filepath.Join(
		cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, "NDQC")), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartSP5D == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartSP5D)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(filepath.Join(
		cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, "S&P5")), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartSP5W == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartSP5W)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(filepath.Join(
		cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, "S&P5")), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartNYCD == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartNYCD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(
		filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, "NYC")), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartNYCW == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartNYCW)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(
		filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, "NYC")), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartDJIAD == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartDJIAD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(
		filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, "DJIA")), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartDJIAW == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartDJIAW)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(
		filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, "DJIA")), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartRUSD == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartRUSD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(
		filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_D", t.Date, "RUS")), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.JsonChartRUSW == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	buffer, err = datautils.FileReaderExtractImage(t.JsonChartRUSW)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(
		filepath.Join(cookiePath, config.StorageNamespaceCharts, fmt.Sprintf("%d_%s_W", t.Date, "RUS")), buffer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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
