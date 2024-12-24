package common

import "fmt"

type Vector2D struct {
	X, Y int
}

type Machine struct {
	ButtonA, ButtonB Vector2D
	Prize            Vector2D
}

func isInteger(f float64) bool {
	return f == float64(int(f))
}

const PriceButtonA = 3
const PriceButtonB = 1

func CalculateCost(machine *Machine) int {
	/*
					Solve system of linear equations via replacing variables.
					Let na = number of A presses, nb = number of B presses
					We have ax*na + bx*nb = px, ay*na + by*nb = py.
			        -> na = (px - bx*nb) / ax
				    Replacing, we get
		   			nb = (py*ax-ay*px)/(ax*by-ay*bx)
					and na = (px-bx*nb) / ax
	*/

	ax := (*machine).ButtonA.X
	ay := (*machine).ButtonA.Y
	bx := (*machine).ButtonB.X
	by := (*machine).ButtonB.Y
	px := (*machine).Prize.X
	py := (*machine).Prize.Y

	nb := float64(py*ax-ay*px) / float64(ax*by-ay*bx)
	na := (float64(px) - float64(bx)*nb) / float64(ax)

	if !isInteger(na) || !isInteger(nb) {
		return 0
	}

	cost := int(na*PriceButtonA) + int(nb*PriceButtonB)
	fmt.Printf("%.3f, %.3f\n", na, nb)
	fmt.Printf("Cost: %d\n", cost)

	return cost
}
