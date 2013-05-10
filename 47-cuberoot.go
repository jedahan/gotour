package main

import "fmt"

func Cbrt(x complex128) complex128 {
	prev_z, z := x, x+0.1

	for prev_z != z {
		prev_z, z = z, z-(z*z*z-x)/(3*z*z)
	}
	return z
}

func main() {
	fmt.Println(Cbrt(2))
}
