package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"config"
	"headerutils"
	"net/http"
	"statistic"
	"strconv"
	"transaction"

	"cloud.google.com/go/datastore"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Statistic(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		statisticGet(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func statisticGet(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	var start int64
	var end int64

	if len(r.URL.Query()) > 0 {

		starts := r.URL.Query().Get("start")
		ends := r.URL.Query().Get("end")

		start, err = strconv.ParseInt(starts, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		end, err = strconv.ParseInt(ends, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	q := datastore.NewQuery(cookie).Namespace(config.NamespaceTransaction)
	q = q.Filter("Order =", "sell")
	q = q.KeysOnly()

	var entities []datastore.PropertyList
	keys, err := client.DatastoreClient.GetAll(client.Context, q, &entities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	orders := make([]*transaction.Order, len(keys))

	err = client.DatastoreClient.GetMulti(client.Context, keys, orders)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	winner := make([]*transaction.Order, 0)
	losser := make([]*transaction.Order, 0)

	for _, o := range orders {

		if o.Date < int(start) || o.Date > int(end) {
			continue
		}

		if o.GainP >= 0.0 {
			winner = append(winner, o)
		} else {
			losser = append(losser, o)
		}

	}

	statistic, err := statistic.NewStatistic(winner, losser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeTemplate(w, "Statistic", statistic, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//type Transaction struct {
//ID   string
//Etag string

//Order string

//Date int

//Symbol string

//Price float64

//Share int

//BuyPoint string

//Revenue float64 `datastore:",omitempty" json:",omitempty"`

//Cost float64 `datastore:",omitempty" json:",omitempty"`

//GainD float64 `datastore:",omitempty" json:",omitempty"`

//GainP float64 `datastore:",omitempty" json:",omitempty"`

//DayHold int `datastore:",omitempty" json:",omitempty"`

//Stage float64

//Note string `datastore:",noindex"`
//}
