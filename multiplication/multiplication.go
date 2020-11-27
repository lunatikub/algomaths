package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"math"
	"time"

	"github.com/lunatikub/lunamath/common"
)

const (
	sz     = 1000
	margin = 10
	x0     = sz / 2 // x ogirin
	y0     = sz / 2 // y origin
	radius = float64(sz / 2)
	pi2    = math.Pi * 2
)

type options struct {
	input     string
	timelapse int
}

func getOptions() *options {
	opts := new(options)
	flag.StringVar(&opts.input, "input", "", "input file in JSON format")
	flag.IntVar(&opts.timelapse, "timelapse", 50, "timelapse between to image in millisecond")
	flag.Parse()
	return opts
}

type fromTo struct {
	Values []float64 `json:"values"`
	From   float64   `json:"from"`
	To     float64   `json:"to"`
	Inc    float64   `json:"inc"`
}

type input struct {
	Table  fromTo `json:"table"`
	Modulo fromTo `json:"modulo"`
}

func get(n, m float64) (float64, float64) {
	x := math.Sin((pi2/m)*n)*radius + x0 + margin
	y := -math.Cos((pi2/m)*n)*radius + y0 + margin
	return x, y
}

var once = true

func multiplication(S *common.SDL, table, modulo float64, timeLapse int) {
	S.Clean()
	for src := 1.0; src < modulo; src++ {
		dst := math.Mod(table*src, modulo)
		xSrc, ySrc := get(src, modulo)
		xDst, yDst := get(dst, modulo)
		S.Line(int(xSrc), int(ySrc), int(xDst), int(yDst), common.Black, 1)
	}
	S.Circle(x0+margin, y0+margin, sz/2, common.Red, 2)
	S.Refresh()
	if once == true {
		time.Sleep(2 * time.Second)
		once = false
	}
	time.Sleep(time.Duration(timeLapse) * time.Millisecond)
}

func foreachModulo(S *common.SDL, I input, t float64, timelapse int) {
	if len(I.Modulo.Values) != 0 {
		for _, m := range I.Modulo.Values {
			multiplication(S, t, m, timelapse)
		}
	} else {
		for m := I.Modulo.From; m <= I.Modulo.To; m += I.Modulo.Inc {
			multiplication(S, t, m, timelapse)
		}
	}
}

func foreachTable(S *common.SDL, I input, timelapse int) {
	if len(I.Table.Values) != 0 {
		for _, t := range I.Table.Values {
			foreachModulo(S, I, t, timelapse)
		}
	} else {
		for t := I.Table.From; t <= I.Table.To; t += I.Table.Inc {
			foreachModulo(S, I, t, timelapse)
		}
	}
}

func main() {
	var I input
	opts := getOptions()
	data, err := ioutil.ReadFile(opts.input)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(data), &I)
	S := common.SDLInit(sz+2*margin, sz+2*margin)

	foreachTable(S, I, opts.timelapse)
	S.Wait()
}
