package device

type Point struct {
	x int
	y int
}

func (m *Point) SetPoint(x, y int) {
	m.x = x
	m.y = y
}

func (m *Point) GetPoint() (int, int) {
	return m.x, m.y
}
func (m *Point) Equals(dest *Point) bool {
	if m.x == dest.x && m.y == dest.y {
		return true
	} else {
		return false
	}
}

func (m *Point) PointEquals(x, y int) bool {
	if m.x == x && m.y == y {
		return true
	} else {
		return false
	}
}

func (m *Point) Initialize(src Point) {
	m.x = src.x
	m.y = src.y
}

func (m *Point) GetVelocity(src Point) (int, int) {
	return m.x - src.x, m.y - src.y
}

func (m *Point) GetOffsetVal(x, y int) (int, int) {
	return m.x + x, m.y + y
}
