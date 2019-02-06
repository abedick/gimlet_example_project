package main

import (
	"github.com/gmbh-micro/gmbh"
)

func main() {
	client, err := gmbh.NewService().Config("../gmbh.yaml")
	if err != nil {
		panic(err)
	}
	client.Verbose().Nonblocking().Start()
}
