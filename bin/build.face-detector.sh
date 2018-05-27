#!/usr/bin/env bash

cd $GOPATH/src/gocv.io/x/gocv
source ./env.sh
cd -
go build -o bin/face-detector detect.go

