package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	prev_z, z := x, x+0.1

	for prev_z != z {
		prev_z, z = z, z-(z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(16))
}
