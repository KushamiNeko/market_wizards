package handler

//import (
//"client"
//"config"
//"headerutils"
//"net/http"
//"statistic"
//"strconv"
//"transaction"

//"cloud.google.com/go/datastore"
//)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func Charts(w http.ResponseWriter, r *http.Request) {

//switch r.Method {
//case http.MethodGet:
//chartsGet(w, r)

//default:
//http.NotFound(w, r)
//}

//}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//func chartsGet(w http.ResponseWriter, r *http.Request) {

//cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
//if err != nil {
//http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
//return
//}

//var start int64
//var end int64

//var threshold float64 = 1.0

////if len(r.URL.Query()) > 0 {

//starts := r.URL.Query().Get("start")
//ends := r.URL.Query().Get("end")

//thresholds := r.URL.Query().Get("threshold")

//if starts != "" {
//start, err = strconv.ParseInt(starts, 10, 64)
//if err != nil {
//http.Error(w, err.Error(), http.StatusInternalServerError)
//return
//}
//}

//if ends != "" {
//end, err = strconv.ParseInt(ends, 10, 64)
//if err != nil {
//http.Error(w, err.Error(), http.StatusInternalServerError)
//return
//}
//}

//if thresholds != "" {
//threshold, err = strconv.ParseFloat(thresholds, 64)
//if err != nil {
//http.Error(w, err.Error(), http.StatusInternalServerError)
//return
//}
//}

////}

//q := datastore.NewQuery(cookie).Namespace(config.NamespaceTransaction)
//q = q.Filter("Order =", "sell")
//q = q.KeysOnly()

//var entities []datastore.PropertyList
//keys, err := client.DatastoreClient.GetAll(client.Context, q, &entities)
//if err != nil {
//http.Error(w, err.Error(), http.StatusInternalServerError)
//return
//}

//orders := make([]*transaction.Order, len(keys))

//err = client.DatastoreClient.GetMulti(client.Context, keys, orders)
//if err != nil {
//http.Error(w, err.Error(), http.StatusInternalServerError)
//return
//}

//winner := make([]*transaction.Order, 0)
//losser := make([]*transaction.Order, 0)

//chartGainDaysHeld := make([][]interface{}, 0)

//chartGainDaysHeld = append(chartGainDaysHeld, []string{
//"DaysHeld",
//"Gain(%)",
//})

//for _, o := range orders {

//if start > 0 && end > 0 && end > start {
//if o.Date < int(start) || o.Date > int(end) {
//continue
//}
//}

////if o.GainP >= statistic.LoserGainThreshold {
//if o.GainP >= threshold {
//winner = append(winner, o)
//} else {
//losser = append(losser, o)
//}

//chartGainDaysHeld = append(chartGainDaysHeld, []float64{
//float64(o.DaysHeld),
//o.GainP,
//})
//}

//statistic, err := statistic.NewStatistic(winner, losser)
//if err != nil {
//http.Error(w, err.Error(), http.StatusInternalServerError)
//return
//}

//statistic.StartDate = starts
//statistic.EndDate = ends
//statistic.LossThresholdP = threshold

//writeTemplate(w, "Statistic", statistic, nil)
//}
