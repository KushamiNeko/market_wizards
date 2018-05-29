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

	//if len(r.URL.Query()) > 0 {

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

	//}

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

	winnerIBD := make([]*bytes.Buffer, 0)
	losserIBD := make([]*bytes.Buffer, 0)

	//chartGeneral := make([][]interface{}, 0)

	//chartGeneral = append(chartGeneral, []interface{}{
	//"DaysHeld",
	//"Gain(%)",
	//})

	filterOrder := make([]*transaction.Order, 0)

	//cookiePath := url.PathEscape(cookie)

	//filepath.Join(cookiePath, config.StorageNamespaceIBDs, fmt.Sprintf("%d_%s", t.Date, t.Symbol)), ibdJsonBuffer)

	for _, o := range orders {

		if start > 0 && end > 0 && end > start {
			if o.Date < int(start) || o.Date > int(end) {
				continue
			}
		}

		filterOrder = append(filterOrder, o)

		//object := filepath.Join(cookiePath, config.StorageNamespaceIBDs, fmt.Sprintf("%d_%s", o.Date, o.Symbol))

		//buffer, err := readStorageObject(object)
		//if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return
		//}

		ikey, exist, err := ibdGetKey(cookie, ibd.IBDCheckupDatastoreGetID(o.Date, o.Symbol))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var ibdDatestore *ibd.IBDCheckupDatastore = nil

		if exist {
			ibdDatestore = new(ibd.IBDCheckupDatastore)

			err = client.DatastoreClient.Get(client.Context, ikey, ibdDatestore)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		//if o.GainP >= statistic.LoserGainThreshold {
		if o.GainP >= threshold {
			winner = append(winner, o)

			if exist {
				winnerIBD = append(winnerIBD, bytes.NewBuffer(ibdDatestore.Data))
			}

			//winnerIBD = append(winnerIBD, buffer)
		} else {
			losser = append(losser, o)

			if exist {
				losserIBD = append(losserIBD, bytes.NewBuffer(ibdDatestore.Data))
			}

			//losserIBD = append(losserIBD, buffer)
		}

		//chartGeneral = append(chartGeneral, []interface{}{
		//o.DaysHeld,
		//o.GainP,
		//})
	}

	stat, err := statistic.NewStatistic(winner, losser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//jsonChartGeneral, err := datautils.JsonB64Encrypt(orders)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//charts := new(statistic.Charts)
	//charts.General = jsonChartGeneral

	stat.StartDate = starts
	stat.EndDate = ends
	stat.LossThresholdP = threshold

	stat.ChartGeneral, err = charts.ChartGeneralNew(filterOrder, winner, losser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeTemplate(w, "Statistic", stat, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func FormatGrpPrice(price int) string {

//grp := math.Floor(float64(price) / statistic.GrpPrice)
//grps := strconv.FormatFloat(grp*statistic.GrpPrice, 'f', -1, 64)
//grpe := strconv.FormatFloat((grp+1)*statistic.GrpPrice, 'f', -1, 64)

//grpk := fmt.Sprintf(statistic.GrpFormat, grps, grpe)

//return grpk
//}

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
