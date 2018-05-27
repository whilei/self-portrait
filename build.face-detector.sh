#!/usr/bin/env bash

cd $GOPATH/src/gocv.io/x/gocv
source ./env.sh
cd -
go build -o face-detector detect.go

