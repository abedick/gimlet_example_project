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

	go func() {
		var err error
		client, err = gmbh.NewService("../gmbh.yaml")
		if err != nil {
			panic(err)
		}
		client.Start()
	}()

	r := mux.NewRouter()
	r.HandleFunc("/", handleIndex)
	r.HandleFunc("/r", handletest)
	r.HandleFunc("/tkn", handletkn)
	log.Fatal(http.ListenAndServe(":3000", r))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cabal Coms Demo\n"))
	w.Write([]byte("Requesting a response from the CORE server...\n"))

	result := client.MakeRequest("demo", "test", "abc123")

	if result.HadError {
		w.Write([]byte("Auth says: " + result.ErrorString + "\n"))
	} else {
		w.Write([]byte("Auth says: " + result.Result + "\n"))
	}
}

func handletest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cabal Coms Demo\n"))
	w.Write([]byte("Requesting a response from the CORE server...\n"))

	result := client.MakeRequest("demo", "two", "abc123")

	if result.HadError {
		w.Write([]byte("Auth says: " + result.ErrorString + "\n"))
	} else {
		w.Write([]byte("Auth says: " + result.Result + "\n"))
	}
}
func handletkn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cabal Coms Demo\n"))
	w.Write([]byte("Requesting a response from the CORE server...\n"))

	result := client.MakeRequest("demo", "tkn", realip.FromRequest(r))

	if result.HadError {
		w.Write([]byte("Auth says: " + result.ErrorString + "\n"))
	} else {
		w.Write([]byte("Auth says: " + result.Result + "\n"))
	}
}
