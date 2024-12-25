package aoc

import "github.com/williamtriinh/advent-of-code/pkg/imath"

var (
	Up    = imath.Vec2{X: 0, Y: -1}
	Right = imath.Vec2{X: 1, Y: 0}
	Down  = imath.Vec2{X: 0, Y: 1}
	Left  = imath.Vec2{X: -1, Y: 0}
)

var Directions = [4]imath.Vec2{Up, Right, Down, Left}
