package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"charts"
	"client"
	"config"
	"headerutils"
	"ibd"
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

	var threshold float64 = 1.0

	starts := r.URL.Query().Get("start")
	ends := r.URL.Query().Get("end")

	thresholds := r.URL.Query().Get("threshold")

	if starts != "" {
		start, err = strconv.ParseInt(starts, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if ends != "" {
		end, err = strconv.ParseInt(ends, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if thresholds != "" {
		threshold, err = strconv.ParseFloat(thresholds, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// get Transaction Orders from datastore

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

	// get IBD Checkups from datastore

	q = datastore.NewQuery(cookie).Namespace(config.NamespaceIBD)
	q = q.KeysOnly()

	keys, err = client.DatastoreClient.GetAll(client.Context, q, &entities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ibdDatastore := make([]*ibd.IBDCheckupDatastore, len(keys))

	err = client.DatastoreClient.GetMulti(client.Context, keys, ibdDatastore)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	winner := make([]*transaction.Order, 0)
	losser := make([]*transaction.Order, 0)

	winnerIBD := make([]*bytes.Buffer, 0)
	losserIBD := make([]*bytes.Buffer, 0)

	filterOrder := make([]*transaction.Order, 0)

	for _, o := range orders {

		if start > 0 && end > 0 && end > start {
			if o.Date < int(start) || o.Date > int(end) {
				continue
			}
		}

		filterOrder = append(filterOrder, o)

		var ibdCheckup *ibd.IBDCheckupDatastore = nil

		for _, c := range ibdDatastore {
			if c.ID == ibd.IBDCheckupDatastoreGetID(o.Date, o.Symbol) {
				ibdCheckup = c
				break
			}
		}

		if o.GainP >= threshold {
			winner = append(winner, o)

			if ibdCheckup != nil {
				winnerIBD = append(winnerIBD, bytes.NewBuffer(ibdCheckup.Data))
			}

		} else {
			losser = append(losser, o)

			if ibdCheckup != nil {
				losserIBD = append(losserIBD, bytes.NewBuffer(ibdCheckup.Data))
			}

		}
	}

	stat, err := statistic.NewStatistic(winner, losser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stat.StartDate = starts
	stat.EndDate = ends
	stat.LossThresholdP = threshold

	stat.ChartGeneral, err = charts.ChartGeneralNew(filterOrder, winner, losser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stat.ChartIBD, err = charts.ChartIBDNew(filterOrder, winnerIBD, losserIBD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeTemplate(w, "Statistic", stat, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
