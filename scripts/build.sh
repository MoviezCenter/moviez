#!/usr/bin/bash

go get .

GOARCH=amd64 GOOS=linux go build -o bin/application .
