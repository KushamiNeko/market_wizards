package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"fmt"
	"headerutils"
	"net/http"
	"strconv"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Caculator(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		caculatorGet(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func caculatorGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	var price float64

	prices := r.URL.Query().Get("Price")

	if prices != "" {
		price, err = strconv.ParseFloat(prices, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(price)

	//writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
