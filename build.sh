#!/bin/bash

go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/mux
go get github.com/tomasen/realip

cd managed_services/
cd gmbh-webserver && make
cd ../gmbh-demo && make
cd ../gmbh-fail && make

cd ../../remote_services
cd gmbh-tester && make

