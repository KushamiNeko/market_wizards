package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"config"
	"context"
	"datautils"
	"headerutils"
	"net/http"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func Eperience(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		experienceGet(w, r)

	//case http.MethodPut:
	//experiencePut(w, r)

	case http.MethodPost:
		experiencePost(w, r)

	default:
		http.NotFound(w, r)
	}

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func experienceGet(w http.ResponseWriter, r *http.Request) {

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	//writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func experiencePost(w http.ResponseWriter, r *http.Request) {

	cookie, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	objectRequest := new(datautils.DateStorage)

	err = datautils.JsonRequestBodyDecode(r, objectRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//chartImg, err := datautils.FileReaderExtractImage(objectRequest.Data)
	//if err != nil {
	//http.Error(w, fmt.Sprintf("Charts Study: %s\n", err.Error()), http.StatusBadRequest)
	//return
	//}

	//chartDatastore := datautils.DateSymbolStorageNewBytes(objectRequest.Date, objectRequest.Symbol, chartImg.Bytes())

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

	_, err := headerutils.GetCookie(r, headerutils.CookieName)
	if err != nil {
		http.Redirect(w, r, Root, http.StatusTemporaryRedirect)
		return
	}

	//writeTemplate(w, "Action", nil, nil)
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
