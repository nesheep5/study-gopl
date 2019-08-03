#!/bin/sh
go build ./lissajous.go
./lissajous > out.gif
open out.gif
