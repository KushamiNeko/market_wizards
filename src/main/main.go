package main

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"client"
	"config"
	"handler"
	"log"
	"net/http"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	//TestFile = "test_files/ibd_checkup_spar.html"
	//TestFile = "test_files/ibd_checkup_cade.html"
	//TestFile = "test_files/ibd_checkup_extr.html"
	//TestFile = "test_files/ibd_checkup_ebsb.html"

	//TestFile = "test_files/ibd_checkup_cacc.html"
	TestFile = "test_files/20180601_MarketSmith_SCVL_W.html"
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
	//buffer := new(bytes.Buffer)

	//f, err := os.Open(TestFile)
	//if err != nil {
	//panic(err)
	//}
	//defer f.Close()

	//io.Copy(buffer, f)

	//m, err := marketsmith.Parse(buffer)
	//if err != nil {
	//panic(err)
	//}

	//fmt.Println(m)

	//return

	mux := http.NewServeMux()

	mux.HandleFunc("/login", handler.Login)

	mux.HandleFunc("/user", handler.User)

	mux.HandleFunc("/action", handler.Action)

	mux.HandleFunc("/transaction", handler.Transaction)

	mux.HandleFunc("/statistic", handler.Statistic)

	mux.HandleFunc("/resource/", handler.Resource)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
