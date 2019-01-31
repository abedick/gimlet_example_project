#!/bin/bash

go get github.com/dgrijalva/jwt-go
go get github.com/gorilla/mux
go get github.com/tomasen/realip

cd services/

cd gmbh-webserver && make

cd ../

cd gmbh-demo && make