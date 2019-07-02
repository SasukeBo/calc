// sqrt.go

package simplemath

import "math"

// Sqrt return int math.Sqrt value
func Sqrt(i int) float64 {
	v := math.Sqrt(float64(i))
	return v
}
