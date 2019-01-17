package device

import "MDIIC/device"

type MousePos struct {
	device.Point
}

func (m *MousePos) GetVelocity(src MousePos) (int, int) {
	return m.X - src.X, m.Y - src.Y
}

func (m *MousePos) GetOffsetVal(X, Y int) (int, int) {
	return m.X + X, m.Y + Y
}

type Mouse struct {
	device.Point
	MoveX int
	MoveY int
}
