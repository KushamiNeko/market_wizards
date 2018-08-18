package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"charts"
	"client"
	"config"
	"context"
	"datautils"
	"encoding/base64"
	"headerutils"
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
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if ends != "" {
		end, err = strconv.ParseInt(ends, 10, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if thresholds != "" {
		threshold, err = strconv.ParseFloat(thresholds, 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	collection := client.MongoClient.Database(config.NamespaceTransaction).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("order", "sell"),
	)

	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	sellOrders := make([]*transaction.Close, 0)

	for cursor.Next(context.Background()) {

		t := new(transaction.Close)

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

	orders := make([]*transaction.Trade, len(sellOrders))

	winners := make([]*transaction.Trade, 0)
	losers := make([]*transaction.Trade, 0)

	winnersIBD := make([]*bytes.Buffer, 0)
	losersIBD := make([]*bytes.Buffer, 0)

	winnersMS := make([]*bytes.Buffer, 0)
	losersMS := make([]*bytes.Buffer, 0)

	for i, sellOrder := range sellOrders {

		collection := client.MongoClient.Database(config.NamespaceTransaction).Collection(cookie)

		filter := bson.NewDocument(
			bson.EC.Interface("order", "buy"),
			bson.EC.Interface("date", sellOrder.DateOfPurchase),
			bson.EC.Interface("symbol", sellOrder.Symbol),
		)

		buyOrder := new(transaction.Open)

		err = collection.FindOne(context.Background(), filter).Decode(buyOrder)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		o := new(transaction.Trade)
		o.Open = buyOrder
		o.Close = sellOrder

		orders[i] = o

		collection = client.MongoClient.Database(config.NamespaceIBD).Collection(cookie)

		filter = bson.NewDocument(
			bson.EC.Interface("date", sellOrder.DateOfPurchase),
			bson.EC.Interface("symbol", sellOrder.Symbol),
		)

		ibdCheckup := new(datautils.DateSymbolStorage)

		ibdErr := collection.FindOne(context.Background(), filter).Decode(ibdCheckup)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		collection = client.MongoClient.Database(config.NamespaceMarketSmith).Collection(cookie)

		filter = bson.NewDocument(
			bson.EC.Interface("date", sellOrder.DateOfPurchase),
			bson.EC.Interface("symbol", sellOrder.Symbol),
		)

		ms := new(datautils.DateSymbolStorage)

		msErr := collection.FindOne(context.Background(), filter).Decode(ms)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if sellOrder.GainP >= threshold {
			t := new(transaction.Trade)
			t.Open = buyOrder
			t.Close = sellOrder

			winners = append(winners, t)

			if ibdErr == mongo.ErrNoDocuments {

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

				winnersIBD = append(winnersIBD, bytes.NewBuffer(data))
			}

			if msErr == mongo.ErrNoDocuments {

			} else {
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				data, err := base64.StdEncoding.DecodeString(ms.Data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				winnersMS = append(winnersMS, bytes.NewBuffer(data))
			}

		} else {
			t := new(transaction.Trade)
			t.Open = buyOrder
			t.Close = sellOrder

			losers = append(losers, t)

			if ibdErr == mongo.ErrNoDocuments {

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

				losersIBD = append(losersIBD, bytes.NewBuffer(data))
			}

			if msErr == mongo.ErrNoDocuments {

			} else {
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				data, err := base64.StdEncoding.DecodeString(ms.Data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				losersMS = append(losersMS, bytes.NewBuffer(data))
			}

		}
	}

	stat, err := statistic.NewStatistic(winners, losers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stat.StartDate = starts
	stat.EndDate = ends
	stat.LossThresholdP = threshold

	stat.ChartGeneral, err = charts.ChartGeneralNew(orders, winners, losers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stat.ChartIBD, err = charts.ChartIBDNew(orders, winnersIBD, losersIBD)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stat.ChartMarketSmith, err = charts.ChartMarketSmithNew(orders, winnersMS, losersMS)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeTemplate(w, "Statistic", stat, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
