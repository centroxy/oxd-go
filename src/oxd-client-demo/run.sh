#!/bin/bash
export GOPATH=`realpath $(pwd)/../..`
go get
go run main.go