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
	"path/filepath"

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

func storagePath(userID, objectID string) string {
	return filepath.Join(url.PathEscape(userID), url.PathEscape(objectID))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func writeStorageObject(object string, buffer io.Reader) error {

	storageBucket := client.StorageClient.Bucket(config.ProjectBucket)

	storageObject := storageBucket.Object(object)

	storageWriter := storageObject.NewWriter(client.Context)

	_, err := io.Copy(storageWriter, buffer)
	if err != nil {
		return err
	}

	storageWriter.Close()

	return nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
