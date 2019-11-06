package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ima := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		imay := make([]uint8, dy)
		ima[x] = imay
		for y := 0; y < dy; y++ {
			ima[x][y]=uint8(x^y)
		}
	}
	return ima
}

func main() {
	pic.Show(Pic)
}

