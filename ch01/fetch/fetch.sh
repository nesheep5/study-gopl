#!/bin/sh

go build fetch.go
./fetch http://gopl.io > out.html
open out.html
