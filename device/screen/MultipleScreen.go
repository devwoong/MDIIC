package device

import "MDIIC/device"

type MultipleScreen struct {
	Main              Screen
	SubScreens        map[string]Screen
	WeightRateScreens map[string]device.Point
}
