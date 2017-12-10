package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"headerutils"
	"net/http"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Transaction(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		transactionGet(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func transactionGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "TransactionNew", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func transactionPost(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func transactionPut(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func transactionDelete(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
