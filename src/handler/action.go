package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"headerutils"
	"net/http"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Action(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		actionGet(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func actionGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
