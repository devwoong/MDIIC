package device

import "MDIIC/device"

type MultipleScreen struct {
	Main              Screen
	SubScreens        map[string]Screen
	WeightRateScreens map[string]device.Point
}

func (m *MultipleScreen) SetWeight() {
	m.WeightRateScreens = make(map[string]device.Point)
	for i, sub := range m.SubScreens {
		widthWeight := (m.Main.width / sub.width)
		heightWeight := (m.Main.height / sub.height)
		weight := device.Point{}
		weight.SetPoint(widthWeight, heightWeight)
		m.WeightRateScreens[i] = weight
	}
}
