package utils

type Direction [2]int

type Position struct {
	Row, Col int
}

type DirectedPosition struct {
	Pos Position
	Dir Direction
}

func Abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
