package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"charts"
	"client"
	"config"
	"context"
	"encoding/base64"
	"headerutils"
	"ibd"
	"net/http"
	"statistic"
	"strconv"
	"transaction"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
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

	collection := client.MongoClient.Database(config.NamespaceTransaction).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("order", "sell"),
		//bson.EC.ArrayFromElements(
		//"$and",
		//bson.VC.DocumentFromElements(
		//bson.EC.SubDocumentFromElements("date",
		//bson.EC.Interface("$gte", start),
		//),
		//),
		//bson.VC.DocumentFromElements(
		//bson.EC.SubDocumentFromElements("date",
		//bson.EC.Interface("$lte", end),
		//),
		//),
		//),
	)

	//t := new(transaction.Order)

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	sellOrders := make([]*transaction.SellOrder, 0)

	for cursor.Next(context.Background()) {

		t := new(transaction.SellOrder)

		err := cursor.Decode(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if start > 0 && end > 0 && end > start {
			if t.Date < int(start) || t.Date > int(end) {
				continue
			}
		}

		sellOrders = append(sellOrders, t)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

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

	// get IBD Checkups from datastore

	//q = datastore.NewQuery(cookie).Namespace(config.NamespaceIBD)
	//q = q.KeysOnly()

	//keys, err = client.DatastoreClient.GetAll(client.Context, q, &entities)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	//ibdDatastore := make([]*ibd.IBDCheckupDatastore, len(keys))

	//err = client.DatastoreClient.GetMulti(client.Context, keys, ibdDatastore)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}

	orders := make([]*transaction.Transaction, len(sellOrders))

	winner := make([]*transaction.Transaction, 0)
	losser := make([]*transaction.Transaction, 0)

	winnerIBD := make([]*bytes.Buffer, 0)
	losserIBD := make([]*bytes.Buffer, 0)

	//filterOrder := make([]*transaction.Order, 0)

	for i, sellOrder := range sellOrders {

		//if start > 0 && end > 0 && end > start {
		//if o.Date < int(start) || o.Date > int(end) {
		//continue
		//}
		//}

		//filterOrder = append(filterOrder, o)

		//var ibdCheckup *ibd.IBDCheckupDatastore = nil

		//for _, c := range ibdDatastore {
		//if c.ID == ibd.IBDCheckupDatastoreGetID(o.Date, o.Symbol) {
		//ibdCheckup = c
		//break
		//}
		//}

		collection := client.MongoClient.Database(config.NamespaceTransaction).Collection(cookie)

		filter := bson.NewDocument(
			bson.EC.Interface("order", "buy"),
			bson.EC.Interface("date", sellOrder.DateOfPurchase),
			bson.EC.Interface("symbol", sellOrder.Symbol),
		)

		buyOrder := new(transaction.BuyOrder)

		err = collection.FindOne(context.Background(), filter).Decode(buyOrder)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		o := new(transaction.Transaction)
		o.Buy = buyOrder
		o.Sell = sellOrder

		orders[i] = o

		collection = client.MongoClient.Database(config.NamespaceIBD).Collection(cookie)

		filter = bson.NewDocument(
			bson.EC.Interface("id", ibd.IBDCheckupDatastoreGetID(sellOrder.DateOfPurchase, sellOrder.Symbol)),
		)

		ibdCheckup := new(ibd.IBDCheckupDatastore)

		err = collection.FindOne(context.Background(), filter).Decode(ibdCheckup)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if sellOrder.GainP >= threshold {
			t := new(transaction.Transaction)
			t.Buy = buyOrder
			t.Sell = sellOrder

			winner = append(winner, t)

			if err == mongo.ErrNoDocuments {
				continue
			} else {
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				data, err := base64.StdEncoding.DecodeString(ibdCheckup.Data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				winnerIBD = append(winnerIBD, bytes.NewBuffer(data))
			}

			//if ibdCheckup != nil {
			//winnerIBD = append(winnerIBD, bytes.NewBuffer(ibdCheckup.Data))
			//}

		} else {
			t := new(transaction.Transaction)
			t.Buy = buyOrder
			t.Sell = sellOrder

			losser = append(losser, t)

			if err == mongo.ErrNoDocuments {
				continue
			} else {
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				data, err := base64.StdEncoding.DecodeString(ibdCheckup.Data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				losserIBD = append(losserIBD, bytes.NewBuffer(data))
			}

			//if ibdCheckup != nil {
			//losserIBD = append(losserIBD, bytes.NewBuffer(ibdCheckup.Data))
			//}

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

	stat.ChartGeneral, err = charts.ChartGeneralNew(orders, winner, losser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stat.ChartIBD, err = charts.ChartIBDNew(orders, winnerIBD, losserIBD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeTemplate(w, "Statistic", stat, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
