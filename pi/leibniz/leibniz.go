package main

import "fmt"

func leibniz(iteration int) float64 {
	n := 0
	sum := 0.0
	numerator := 1.0
	denominator := 1.0

	for {
		sum += numerator / denominator
		numerator *= -1
		denominator += 2
		n++
		if n == iteration {
			break
		}
	}
	return sum * 4.0
}

func main() {
	iterations := []int{100, 1000, 10000, 100000, 10000000, 100000000}
	for _, i := range iterations {
		fmt.Println(i, leibniz(i))
	}
}
