package main

import (
	"log"
	"net/http"

	"github.com/gmbh-micro/gmbh"
	"github.com/gorilla/mux"
	"github.com/tomasen/realip"
)

var client *gmbh.Client

func main() {

	var err error
	client, err = gmbh.NewService().Config("../gmbh.yaml")
	if err != nil {
		panic(err)
	}
	client.Verbose().Nonblocking().Start()

	r := mux.NewRouter()
	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/r", handletest)
	r.HandleFunc("/about", handleAbout)
	r.HandleFunc("/info", handleInfo)
	r.HandleFunc("/tkn", handletkn)
	log.Fatal(http.ListenAndServe(":3000", r))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cabal Coms Demo\n"))
	w.Write([]byte("Requesting a response from the CORE server...\n"))

	result, _ := client.MakeRequest("demo", "test", "abc123")

	if result.HadError {
		w.Write([]byte("Auth says: " + result.ErrorString + "\n"))
	} else {
		w.Write([]byte("Auth says: " + result.Result + "\n"))
	}
}

func handletest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cabal Coms Demo\n"))
	w.Write([]byte("Requesting a response from the CORE server...\n"))

	result, err := client.MakeRequest("demo", "two", "abc123")
	if err != nil {
		panic(err)
	}

	if result.HadError {
		w.Write([]byte("Auth says: " + result.ErrorString + "\n"))
	} else {
		w.Write([]byte("Auth says: " + result.Result + "\n"))
	}
}

func handleAbout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("gmbh-cabal demo\n"))
	w.Write([]byte("Requesting a response from remote service...\n"))

	result, err := client.MakeRequest("tester", "about", "")
	if err != nil {
		panic(err)
	}

	if result.HadError {
		w.Write([]byte("About says: " + result.ErrorString + "\n"))
	} else {
		w.Write([]byte("About says: " + result.Result + "\n"))
	}
}

func handleInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("gmbh-cabal demo\n"))
	w.Write([]byte("Requesting a response from remote service...\n"))

	result, err := client.MakeRequest("info", "info", "")
	if err != nil {
		panic(err)
	}

	if result.HadError {
		w.Write([]byte("About says: " + result.ErrorString + "\n"))
	} else {
		w.Write([]byte("About says: " + result.Result + "\n"))
	}
}

func handletkn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cabal Coms Demo\n"))
	w.Write([]byte("Requesting a response from the CORE server...\n"))

	result, err := client.MakeRequest("demo", "tkn", realip.FromRequest(r))
	if err != nil {
		panic(err)
	}

	if result.HadError {
		w.Write([]byte("Auth says: " + result.ErrorString + "\n"))
	} else {
		w.Write([]byte("Auth says: " + result.Result + "\n"))
	}
}
