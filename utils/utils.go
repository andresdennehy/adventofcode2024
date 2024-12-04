package utils

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
