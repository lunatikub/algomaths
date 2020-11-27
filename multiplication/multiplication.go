package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"math"
	"time"

	"github.com/lunatikub/lunamath/common"
)

// constant using for the multiplication
const (
	sz     = 1000
	margin = 10
	x0     = sz / 2          // x ogirin
	y0     = sz / 2          // y origin
	R      = float64(sz / 2) // radius
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

type v struct {
	Values []float64 `json:"values"`
	From   float64   `json:"from"`
	To     float64   `json:"to"`
	Step   float64   `json:"step"`
}

type input struct {
	Table  v `json:"table"`
	Modulo v `json:"modulo"`
}

// Get the coordinates of the point depending on the angle
func get(alpha, N float64) (float64, float64) {
	x := math.Sin(N*alpha) * R
	y := math.Cos(N*alpha) * R
	return x, y
}

// Convert theorical coordinates to SDL coordinates
func convert(x, y float64) (float64, float64) {
	return x + x0 + margin, -y + y0 + margin
}

// Draw a line between 2 points
func line(S *common.SDL, alpha, N, T, M float64) {
	R := math.Mod(T*N, M)
	x1, y1 := convert(get(alpha, N))
	x2, y2 := convert(get(alpha, R))
	S.Line(int(x1), int(y1), int(x2), int(y2), common.Black, 1)
}

// T: table
// M: modulo
func multiplication(S *common.SDL, T, M float64, timeLapse int) {
	S.Clean()
	alpha := (2 * math.Pi) / M
	for N := 1.0; N < M; N++ {
		line(S, alpha, N, T, M)
	}
	S.Circle(x0+margin, y0+margin, sz/2, common.Red, 2)
	S.Refresh()
	time.Sleep(time.Duration(timeLapse) * time.Millisecond)
}

func foreachModulo(S *common.SDL, I input, t float64, timelapse int) {
	if len(I.Modulo.Values) != 0 {
		for _, m := range I.Modulo.Values {
			multiplication(S, t, m, timelapse)
		}
	} else {
		for m := I.Modulo.From; m <= I.Modulo.To; m += I.Modulo.Step {
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
		for t := I.Table.From; t <= I.Table.To; t += I.Table.Step {
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
