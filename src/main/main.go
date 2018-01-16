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
	//TestFile = "resource/ibd_checkup_spar.html"
	//TestFile = "resource/ibd_checkup_cacc.html"
	//TestFile = "resource/ibd_checkup_cade.html"
	TestFile = "resource/ibd_checkup_extr.html"
	//TestFile = "resource/ibd_checkup_ebsb.html"
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

	//info, err := ibd.Parse(buffer)
	//if err != nil {
	//panic(err)
	//}

	//fmt.Println(info)

	//return

	mux := http.NewServeMux()

	mux.HandleFunc("/login", handler.Login)

	mux.HandleFunc("/user", handler.User)

	mux.HandleFunc("/action", handler.Action)

	mux.HandleFunc("/transaction", handler.Transaction)

	mux.HandleFunc("/statistic", handler.Statistic)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
