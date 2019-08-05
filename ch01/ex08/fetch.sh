#!/bin/sh

go build -o bin/fetch fetch.go
./bin/fetch gopl.io > out.html
open out.html
