package common

import "math"

// Radian convert degrees to radian
func Radian(angle float64) float64 {
	return angle * math.Pi / 180
}

// GetIntersections Get the next intersection between a circle
// centered at {0,0} with radius = R and a line with
// equation Ax + By + C = 0
func GetIntersections(A, B, C, R float64) (float64, float64, float64, float64) {
	A2 := math.Pow(A, 2)
	B2 := math.Pow(B, 2)
	R2 := math.Pow(R, 2)
	C2 := math.Pow(C, 2)

	x0 := -((A * C) / (A2 + B2))
	y0 := -((B * C) / (A2 + B2))

	d := math.Sqrt(R2 - C2/(A2+B2))
	m := math.Sqrt(math.Pow(d, 2) / (A2 + B2))

	x1 := x0 + B*m
	y1 := y0 - A*m

	x2 := x0 - B*m
	y2 := y0 + A*m

	return x1, y1, x2, y2
}
