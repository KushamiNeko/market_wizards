package main

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"client"
	"config"
	"handler"
	"ibd"
	"io"
	"log"
	"net/http"
	"os"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	//TestFile = "test_files/ibd_checkup_spar.html"
	//TestFile = "test_files/ibd_checkup_cade.html"
	//TestFile = "test_files/ibd_checkup_extr.html"
	//TestFile = "test_files/ibd_checkup_ebsb.html"

	TestFile = "test_files/ibd_checkup_cacc.html"
	//TestFile = "test_files/ibd_checkup_sgh.html"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {

	// config Init before anything
	config.Init()

	client.Init()

}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	buffer := new(bytes.Buffer)

	f, err := os.Open(TestFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(buffer, f)

	//ibd, err := ibd.Parse(buffer)
	_, err = ibd.Parse(buffer)
	if err != nil {
		panic(err)
	}

	//fmt.Println(ibd)

	//jsonBuffer, err := json.Marshal(ibd)
	//if err != nil {
	//panic(err)
	//}

	//storageBucket := client.StorageClient.Bucket(config.ProjectBucket)

	//storageObject := storageBucket.Object("test")

	//storageWriter := storageObject.NewWriter(client.Context)

	//storageWriter.Write(jsonBuffer)

	//storageWriter.Close()

	return

	mux := http.NewServeMux()

	mux.HandleFunc("/login", handler.Login)

	mux.HandleFunc("/user", handler.User)

	mux.HandleFunc("/action", handler.Action)

	mux.HandleFunc("/transaction", handler.Transaction)

	mux.HandleFunc("/statistic", handler.Statistic)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
