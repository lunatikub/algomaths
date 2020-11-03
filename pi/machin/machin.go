package main

import (
	"fmt"
	"math"
)

func johnMachin() float64 {
	n1 := 4 * math.Atan2(1, 5)
	n2 := math.Atan2(1, 239)
	return 4 * (n1 - n2)
}

func main() {
	fmt.Println(johnMachin())
}
