package main

import (
	"log"
	"net/http"

	"github.com/gimlet-gmbh/gmbh"
	"github.com/gorilla/mux"
	"github.com/tomasen/realip"
)

var gimlet *gmbh.Gimlet

func main() {

	go func() {
		var err error
		gimlet, err = gmbh.NewService("../webserver.yaml")
		if err != nil {
			panic(err)
		}
		gimlet.Start()
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

	result := gimlet.MakeRequest("demo", "test", "abc123")

	if result.HadError {
		w.Write([]byte("Auth says: " + result.ErrorString + "\n"))
	} else {
		w.Write([]byte("Auth says: " + result.Result + "\n"))
	}
}

func handletest(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cabal Coms Demo\n"))
	w.Write([]byte("Requesting a response from the CORE server...\n"))

	result := gimlet.MakeRequest("demo", "two", "abc123")

	if result.HadError {
		w.Write([]byte("Auth says: " + result.ErrorString + "\n"))
	} else {
		w.Write([]byte("Auth says: " + result.Result + "\n"))
	}
}
func handletkn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cabal Coms Demo\n"))
	w.Write([]byte("Requesting a response from the CORE server...\n"))

	result := gimlet.MakeRequest("demo", "tkn", realip.FromRequest(r))

	if result.HadError {
		w.Write([]byte("Auth says: " + result.ErrorString + "\n"))
	} else {
		w.Write([]byte("Auth says: " + result.Result + "\n"))
	}
}
