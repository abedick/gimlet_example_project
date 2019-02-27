package main

import (
	"github.com/gmbh-micro/gmbh"
)

func main() {
	runtime := gmbh.SetRuntime(gmbh.RuntimeOptions{Blocking: false, Verbose: true})
	client, err := gmbh.NewClient("./sf.toml", runtime)
	if err != nil {
		panic(err)
	}
	client.Start()
}
