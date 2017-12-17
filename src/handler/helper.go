package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"client"
	"config"
	"io"
	"minify"
	"net/http"
	"net/url"

	"cloud.google.com/go/datastore"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func emailExist(email string) ([]*datastore.Key, bool, error) {

	q := datastore.NewQuery(config.KindUser).Namespace(config.NamespaceUser)
	q = q.Filter("Email =", email)
	q = q.KeysOnly()

	var entities []datastore.PropertyList
	keys, err := client.DatastoreClient.GetAll(client.Context, q, &entities)
	if err != nil {
		return nil, false, err
	}

	if len(keys) > 0 {
		return keys, true, nil
	}

	return nil, false, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func uidExist(uid string) ([]*datastore.Key, bool, error) {

	q := datastore.NewQuery(config.KindUser).Namespace(config.NamespaceUser)
	q = q.Filter("UID =", uid)
	q = q.KeysOnly()

	var entities []datastore.PropertyList
	keys, err := client.DatastoreClient.GetAll(client.Context, q, &entities)
	if err != nil {
		return nil, false, err
	}

	if len(keys) > 0 {
		return keys, true, nil
	}

	return nil, false, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func writeTemplate(w http.ResponseWriter, template string, data interface{}, cb func()) {
	buffer := new(bytes.Buffer)

	err := templates.ExecuteTemplate(
		buffer,
		template,
		data,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if cb != nil {
		cb()
	}

	w.Write(minify.Minify(buffer.Bytes()))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func writeStorageObject(object string, content string) error {

	storageBucket := client.StorageClient.Bucket(config.ProjectBucket)
	storageObject := storageBucket.Object(url.PathEscape(object))

	storageWriter := storageObject.NewWriter(client.Context)

	_, err := io.WriteString(storageWriter, content)
	if err != nil {
		return err
	}

	storageWriter.Close()

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
