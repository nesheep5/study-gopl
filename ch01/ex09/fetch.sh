#!/bin/sh

go build -o bin/fetch fetch.go
./bin/fetch http://gopl.io 
