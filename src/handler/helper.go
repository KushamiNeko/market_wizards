package handler

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"client"
	"config"
	"fmt"
	"io"
	"minify"
	"net/http"

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

func ibdGetKey(kind, id string) (*datastore.Key, bool, error) {

	q := datastore.NewQuery(kind).Namespace(config.NamespaceIBD)
	q = q.Filter("ID =", id)
	q = q.KeysOnly()

	var entities []datastore.PropertyList
	keys, err := client.DatastoreClient.GetAll(client.Context, q, &entities)
	if err != nil {
		return nil, false, err
	}

	if len(keys) > 1 {
		//errorlog.ArchitectureLogicalError("More than one entity for each IBD ID\n")
		return nil, false, fmt.Errorf("More than one entity for each IBD ID\n")
	}

	if len(keys) > 0 {
		return keys[0], true, nil
	}

	key := datastore.IncompleteKey(kind, nil)
	key.Namespace = config.NamespaceIBD

	return key, false, nil
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

//func storagePath(userID, objectID string) string {
//return filepath.Join(url.PathEscape(userID), url.PathEscape(objectID))
//}

//func storagePath(userID, folder, objectID string) string {
//return filepath.Join(userID, folder, objectID)
//}

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

func readStorageObject(object string) (*bytes.Buffer, error) {

	storageBucket := client.StorageClient.Bucket(config.ProjectBucket)
	storageObject := storageBucket.Object(object)

	storageReader, err := storageObject.NewReader(client.Context)
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)

	_, err = io.Copy(buffer, storageReader)
	if err != nil {
		return nil, err
	}

	storageReader.Close()

	return buffer, nil

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
