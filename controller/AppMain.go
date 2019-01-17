package controller

import (
	"MDIIC/common"
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
	SendMessage     chan common.Message
	RecvMessage     chan common.Message
	isServer        bool
}

func (app *appObject) AppMain(isServer bool) {
	// robotgo.ScrollMouse(10, "up")
	// robotgo.MouseClick("left", true)
	// robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)

	currentPos := mouse.MousePos{Point: device.Point{0, 0}}
	prevPos := mouse.MousePos{device.Point{0, 0}}
	app.isServer = isServer

	app.Screen.Main.SetSize(robotgo.GetScreenSize())
	width, _ := app.Screen.Main.GetSize()
	for {
		currentPos.SetPoint(robotgo.GetMousePos())
		if currentPos.Equals(&prevPos.Point) {
			continue
		} else {
			cx, cy := currentPos.GetPoint()
			px, _ := prevPos.GetPoint()
			//event
			// 좌 끝단 도달
			if cx <= 0 && px <= 0 {
				fmt.Printf("좌 끝단 : x: %d y:  %d\n", cx, cy)
			} else if cx >= width && px >= width {
				fmt.Printf("우 끝단 : x: %d y:  %d\n", cx, cy)
				ox, oy := currentPos.GetOffsetVal(currentPos.GetVelocity(prevPos))

				mouseEvent := mouse.Mouse{device.Point{cx, cy}, ox, oy}

				mouseMove := common.Message{}
				mouseMove.Type = common.MSG_MOUSE
				mouseMove.Code = common.MOUSE_MOVE
				mouseMove.IsServer = isServer
				mouseMove.Message = common.ObjectToByte(mouseEvent)
				app.SendMessage <- mouseMove
			}
			prevPos.Initialize(currentPos.Point)

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
		appInstance.SendMessage = make(chan common.Message)
		appInstance.RecvMessage = make(chan common.Message)
		appInstance.Screen = screen.NewMultiScreen()
	}
	return appInstance
}
