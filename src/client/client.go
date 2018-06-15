package client

//////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"config"
	"context"

	"github.com/mongodb/mongo-go-driver/mongo"
	//"golang.org/x/net/context"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	//Context context.Context
	//DatastoreClient *datastore.Client
	//StorageClient   *storage.Client

	MongoClient *mongo.Client
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func Init() {

	var err error

	//Context = context.Background()

	//DatastoreClient, err = datastore.NewClient(Context, config.ProjectID)
	//if err != nil {
	//panic(err)
	//}

	//StorageClient, err = storage.NewClient(Context)
	//if err != nil {
	//panic(err)
	//}

	MongoClient, err = mongo.NewClient(config.MongoURL)
	if err != nil {
		panic(err)
	}

	err = MongoClient.Connect(context.TODO())
	if err != nil {
		panic(err)
	}
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
