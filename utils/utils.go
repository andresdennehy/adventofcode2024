package utils

import "math"

type Direction [2]int

type Position struct {
	Row, Col int
}

type DirectedPosition struct {
	Pos Position
	Dir Direction
}

var UpRightDownLeft = []Direction{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
