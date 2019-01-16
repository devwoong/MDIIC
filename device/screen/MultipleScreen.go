package device

import "MDIIC/device"

type MultipleScreen struct {
	Main              Screen
	SubScreens        map[int]Screen
	WeightRateScreens map[int]device.Point
}

func (m *MultipleScreen) SetWeight() {
	m.WeightRateScreens = make(map[int]device.Point)
	for i, sub := range m.SubScreens {
		widthWeight := (m.Main.width / sub.width)
		heightWeight := (m.Main.height / sub.height)
		weight := device.Point{}
		weight.SetPoint(widthWeight, heightWeight)
		m.WeightRateScreens[i] = weight
	}
}

func NewMultiScreen() MultipleScreen {
	resultScreen := MultipleScreen{}
	resultScreen.SubScreens = make(map[int]Screen)
	resultScreen.WeightRateScreens = make(map[int]device.Point)
	return resultScreen
}
