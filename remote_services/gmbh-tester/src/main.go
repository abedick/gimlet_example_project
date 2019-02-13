package main

import (
	"github.com/gmbh-micro/gmbh"
)

func main() {
	runtime := gmbh.SetRuntime(gmbh.RuntimeOptions{Blocking: true, Verbose: true})
	client, err := gmbh.NewClient("../gmbh.yaml", runtime)
	if err != nil {
		panic(err)
	}

	client.Route("about", handleAbout)
	client.Start()
}

func handleAbout(req gmbh.Request, resp *gmbh.Responder) {
	resp.Result = "About: This is a remote tester service."
	return
}
