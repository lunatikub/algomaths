package main

import (
	"flag"
	"fmt"
	"math"

	common "github.com/lunatikub/AlgoMathAndCo/common"
)

type options struct {
	m1 float64
	m2 float64
}

func getOptions() *options {
	opts := new(options)
	flag.Float64Var(&opts.m1, "m1", 1, "mass m1")
	flag.Float64Var(&opts.m2, "m2", 1, "mass m2")
	flag.Parse()
	return opts
}

func main() {
	v1 := -1.0 // initial velocity of block 1
	v2 := 0.0  // itnitial velocity of block 2

	opts := getOptions()

	A := math.Sqrt(opts.m1)
	B := math.Sqrt(opts.m2)
	R := A
	collisions := 0

	for {
		collisions++ // blocks collision
		C := -(opts.m1*v1 + opts.m2*v2)
		x, y, _, _ := common.GetIntersections(A, B, C, R)
		v1 = x / math.Sqrt(opts.m1)
		v2 = y / math.Sqrt(opts.m2)
		if v2 < 0 {
			collisions++ // wall collision
			v2 = -v2
		}
		if v1 > v2 && v1 > 0 {
			break // no more collision
		}
	}
	fmt.Println("collisions:", collisions)
}
