package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"config"
	"headerutils"
	"net/http"
	"statistic"
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
