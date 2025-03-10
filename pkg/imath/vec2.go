package imath

type Vec2 struct {
	X int
	Y int
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{X: v.X + other.X, Y: v.Y + other.Y}
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
	return Vec2{X: v.X - other.X, Y: v.Y - other.Y}
}

func ManhattanDistanceVec2(a, b Vec2) int {
	return Abs(a.X-b.X) + Abs(a.Y-b.Y)
}
