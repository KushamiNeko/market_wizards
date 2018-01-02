package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"headerutils"
	"net/http"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Statistic(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		return

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func statisticGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	//writeTemplate(w, "TransactionNew", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
