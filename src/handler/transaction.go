package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"config"
	"datautils"
	"headerutils"
	"net/http"
	"transaction"

	"cloud.google.com/go/datastore"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Transaction(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		if len(r.URL.Query()) == 0 {
			transactionNewGet(w, r)
		}

	case http.MethodPost:
		transactionPost(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func transactionNewGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "TransactionNew", nil, nil)
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

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	t := new(transaction.Buy)
	err = datautils.JsonRequestBodyDecode(r, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := client.DatastoreClient.NewTransaction(client.Context)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	key := datastore.IncompleteKey(cookie, nil)
	key.Namespace = config.NamespaceTransaction

	_, err = tx.Put(key, t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = tx.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusPreconditionFailed)
		return
	}

	var content string

	content = datautils.FileReaderExtractContent(t.JsonChartD)
	if content == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(t.ChartD, content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	content = datautils.FileReaderExtractContent(t.JsonChartW)
	if content == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = writeStorageObject(t.ChartW, content)
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
