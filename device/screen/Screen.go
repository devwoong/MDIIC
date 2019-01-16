package device

import "MDIIC/device"

type Screen struct {
	device.Point
	x      int
	y      int
	width  int
	height int
}

func (s *Screen) SetSize(width, height int) {
	s.width = width
	s.height = height
}

func (s *Screen) getSize() (int, int) {
	return s.width, s.height
}

func (s *Screen) getScreen() (int, int, int, int) {
	return s.x, s.y, s.width, s.height
}
