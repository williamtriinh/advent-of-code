package imath

type Vec2 struct {
	X int
	Y int
}

func (v Vec2) Dot(other Vec2) int {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec2) Equals(other Vec2) bool {
	return v.X == other.X && v.Y == other.Y
}

func (v Vec2) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v Vec2) Subtract(other Vec2) Vec2 {
	return Vec2{X: other.X - v.X, Y: other.Y - v.Y}
}
