package main

import (
	"time"

	"github.com/gmbh-micro/gmbh"
)

func main() {

	runtime := gmbh.SetRuntime(gmbh.RuntimeOptions{Blocking: false, Verbose: true})
	client, err := gmbh.NewService("../gmbh.yaml", runtime)
	if err != nil {
		panic(err)
	}

	client.Route("info", handleAbout)

	client.Start()

	time.Sleep(time.Second * 2)
}

func handleAbout(req gmbh.Request, resp *gmbh.Responder) {
	resp.Result = "Info: This is a remote tester service."
	return
}
