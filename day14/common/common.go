package common

type Robot struct {
	Col, Row   int
	DCol, DRow int
}

const SpaceRows = 103
const SpaceCols = 101

func (robot *Robot) Move() {
	robot.Col += robot.DCol
	if robot.Col > SpaceCols-1 {
		robot.Col = robot.Col % SpaceCols
	}
	if robot.Col < 0 {
		robot.Col += SpaceCols
	}
	robot.Row += robot.DRow
	if robot.Row > SpaceRows-1 {
		robot.Row = robot.Row % SpaceRows
	}
	if robot.Row < 0 {
		robot.Row += SpaceRows
	}
}
