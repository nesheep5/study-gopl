#!/bin/sh
go build -o bin/fetchall . 
./bin/fetchall https://golang.org http://gopl.io https://godoc.org
