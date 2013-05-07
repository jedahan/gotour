package main

import "code.google.com/p/go-tour/pic"

const width, height = 640, 640

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, height)
	for y, _ := range pic {
		pic[y] = make([]uint8, width)
		for x, _ := range pic[y] {
			pic[y][x] = uint8(x ^ y)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}
