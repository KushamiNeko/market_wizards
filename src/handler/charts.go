package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"headerutils"
	"net/http"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Charts(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		chartsGet(w, r)

	case http.MethodPost:
		chartsPost(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func chartsGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	//writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func chartsPost(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	//writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
