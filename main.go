package main

import (
	"MDIIC/device"
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {
	// robotgo.ScrollMouse(10, "up")
	// robotgo.MouseClick("left", true)
	// robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)

	currentPos := device.Point{}
	currentPos.SetPoint(0, 0)
	prevPos := device.Point{}
	prevPos.SetPoint(0, 0)
	for {
		currentPos.SetPoint(robotgo.GetMousePos())
		if currentPos.Equals(&prevPos) {
			continue
		} else {
			// x, y := currentPos.GetPoint()
			// fmt.Printf("currentCursol : x: %d y:  %d\n", x, y)
			// /// event
			// vx, vy := currentPos.GetVelocity(prevPos)
			// fmt.Printf("velocity : x: %d y:  %d\n", vx, vy)

			// ox, oy := currentPos.GetOffsetVal(currentPos.GetVelocity(prevPos))
			// fmt.Printf("offset : x: %d y:  %d\n", ox, oy)

			// robotgo.MoveMouse(currentPos.GetOffsetVal(currentPos.GetVelocity(prevPos)))
			// prevPos.SetPoint(robotgo.GetMousePos())
			//event
			prevPos.Initialize(currentPos)

		}
		time.Sleep(time.Millisecond * 50)
	}
}
