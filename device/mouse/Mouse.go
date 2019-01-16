package device

import "MDIIC/device"

type MousePos struct {
	device.Point
	x int
	y int
}

func (m *MousePos) GetVelocity(src MousePos) (int, int) {
	return m.x - src.x, m.y - src.y
}

func (m *MousePos) GetOffsetVal(x, y int) (int, int) {
	return m.x + x, m.y + y
}
