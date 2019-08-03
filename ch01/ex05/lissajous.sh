#!/bin/sh
go build -o ./bin/lissajous lissajous.go
./bin/lissajous > out.gif

# open out.gif
open -a 'Google Chrome' out.gif
