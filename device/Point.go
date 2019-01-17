package device

type Point struct {
	X int
	Y int
}

func (m *Point) SetPoint(X, Y int) {
	m.X = X
	m.Y = Y
}

func (m *Point) GetPoint() (int, int) {
	return m.X, m.Y
}
func (m *Point) Equals(dest *Point) bool {
	if m.X == dest.X && m.Y == dest.Y {
		return true
	} else {
		return false
	}
}

func (m *Point) PointEquals(X, Y int) bool {
	if m.X == X && m.Y == Y {
		return true
	} else {
		return false
	}
}

func (m *Point) Initialize(src Point) {
	m.X = src.X
	m.Y = src.Y
}

func (m *Point) GetVelocitY(src Point) (int, int) {
	return m.X - src.X, m.Y - src.Y
}

func (m *Point) GetOffsetVal(X, Y int) (int, int) {
	return m.X + X, m.Y + Y
}
