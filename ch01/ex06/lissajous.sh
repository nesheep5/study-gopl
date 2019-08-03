#!/bin/sh
go build ./lissajous.go
./lissajous > out.gif

# open out.gif
open -a 'Google Chrome' out.gif
