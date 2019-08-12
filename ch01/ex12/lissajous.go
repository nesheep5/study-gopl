package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	shiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/lissajous", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	c, err := strconv.Atoi(r.URL.Query().Get("cycles"))
	if err != nil {
		fmt.Fprintf(w, "invalid param")
	}
	l := &lissajous{
		cycles:  float64(c),
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}
	l.drawing(w)
}

type lissajous struct {
	cycles  float64
	res     float64
	size    int
	nframes int
	delay   int
}

func (l *lissajous) drawing(out io.Writer) {
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: l.nframes}
	phase := 0.0
	for i := 0; i < l.nframes; i++ {
		rect := image.Rect(0, 0, 2*l.size+1, 2*l.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < l.cycles*2*math.Pi; t += l.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(l.size+int(x*float64(l.size)+0.5),
				l.size+int(y*float64(l.size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, l.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
