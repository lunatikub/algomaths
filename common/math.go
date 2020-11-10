package common

import "math"

// Radian convert degrees to radian
func Radian(angle float64) float64 {
	return angle * math.Pi / 180
}
