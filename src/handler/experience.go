package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"config"
	"context"
	"datautils"
	"headerutils"
	"net/http"
	"strings"

	"github.com/mongodb/mongo-go-driver/bson"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Experience(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		action := r.URL.Query().Get("action")

		if action == "" {
			experienceGet(w, r)
		} else if action == "new" {
			experienceNewGet(w, r)
		} else {
			http.NotFound(w, r)
		}

	case http.MethodPut:
		experiencePut(w, r)

	case http.MethodPost:
		experiencePost(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func experienceGet(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	collection := client.MongoClient.Database(config.NamespaceExperience).Collection(cookie)

	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.Background())

	//items := make([]*datautils.IdStorage, 0)

	items := make([][]string, 0)

	for cursor.Next(context.Background()) {

		t := new(datautils.IdStorage)

		err := cursor.Decode(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := strings.Split(t.Data, "\n")

		items = append(items, data)

		//items = append(items, t)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeTemplate(w, "Experience", items, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func experienceNewGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	writeTemplate(w, "ExperienceNew", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func experiencePost(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	objectRequest := new(datautils.IdStorage)

	err = datautils.JsonRequestBodyDecode(r, objectRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.MongoClient.Database(config.NamespaceExperience).Collection(cookie)

	//filter := bson.NewDocument(
	//bson.EC.Interface("date", objectRequest.Date),
	//)

	//err = collection.FindOne(context.Background(), filter).Decode(&datautils.DateStorage{})
	//if err == nil {
	//_, err = collection.ReplaceOne(context.Background(), filter, chartDatastore)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}
	//} else {
	//if err == mongo.ErrNoDocuments {
	//_, err = collection.InsertOne(context.Background(), chartDatastore)
	//if err != nil {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}
	//} else {
	//http.Error(w, err.Error(), http.StatusInternalServerError)
	//return
	//}
	//}

	_, err = collection.InsertOne(context.Background(), objectRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func experiencePut(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	objectRequest := new(datautils.DateIdStorage)

	err = datautils.JsonRequestBodyDecode(r, objectRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	collection := client.MongoClient.Database(config.NamespaceExperience).Collection(cookie)

	filter := bson.NewDocument(
		bson.EC.Interface("id", objectRequest.Id),
	)

	err = collection.FindOne(context.Background(), filter).Decode(&datautils.DateStorage{})
	if err == nil {
		_, err = collection.ReplaceOne(context.Background(), filter, objectRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		//if err == mongo.ErrNoDocuments {
		//_, err = collection.InsertOne(context.Background(), chartDatastore)
		//if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return
		//}
		//} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
		//}
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
