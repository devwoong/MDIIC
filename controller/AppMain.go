package controller

import (
	"MDIIC/device"
	mouse "MDIIC/device/mouse"
	screen "MDIIC/device/screen"
	"fmt"
	"sync"
	"time"

	"github.com/go-vgo/robotgo"
)

type appObject struct {
	Screen          screen.MultipleScreen
	SendkeyboardMsg chan string
	RecvkeyboardMsg chan string
	SendMouseMsg    chan mouse.Mouse
	RecvMouseMsg    chan mouse.Mouse
	isServer        bool
}

func (app *appObject) AppMain(isServer bool) {
	// robotgo.ScrollMouse(10, "up")
	// robotgo.MouseClick("left", true)
	// robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)

	currentPos := device.Point{}
	currentPos.SetPoint(0, 0)
	prevPos := device.Point{}
	prevPos.SetPoint(0, 0)
	app.isServer = isServer

	app.Screen.Main.SetSize(robotgo.GetScreenSize())
	width, _ := app.Screen.Main.GetSize()
	for {
		currentPos.SetPoint(robotgo.GetMousePos())
		if currentPos.Equals(&prevPos) {
			continue
		} else {
			cx, cy := currentPos.GetPoint()
			px, _ := prevPos.GetPoint()
			// fmt.Printf("currentCursol : x: %d y:  %d\n", x, y)
			// /// event
			// vx, vy := currentPos.GetVelocity(prevPos)
			// fmt.Printf("velocity : x: %d y:  %d\n", vx, vy)

			// ox, oy := currentPos.GetOffsetVal(currentPos.GetVelocity(prevPos))
			// fmt.Printf("offset : x: %d y:  %d\n", ox, oy)

			// robotgo.MoveMouse(currentPos.GetOffsetVal(currentPos.GetVelocity(prevPos)))
			// prevPos.SetPoint(robotgo.GetMousePos())
			//event
			// 좌 끝단 도달
			if cx <= 0 && px <= 0 {
				fmt.Printf("좌 끝단 : x: %d y:  %d\n", cx, cy)
			} else if cx >= width && px >= width {
				fmt.Printf("우 끝단 : x: %d y:  %d\n", cx, cy)

			}
			prevPos.Initialize(currentPos)

		}
		time.Sleep(time.Millisecond * 50)
	}
}

var appInstance *appObject = nil
var mu sync.Mutex

func GetInstance() *appObject {
	mu.Lock()
	defer mu.Unlock()
	if appInstance == nil {
		appInstance = &appObject{}
		appInstance.SendMouseMsg = make(chan mouse.Mouse)
		appInstance.RecvMouseMsg = make(chan mouse.Mouse)
		appInstance.Screen = screen.NewMultiScreen()
	}
	return appInstance
}
