package controller

import (
	"MDIIC/common"
	"MDIIC/device"
	mouse "MDIIC/device/mouse"
	"fmt"

	"github.com/go-vgo/robotgo"
)

type MouseEvent struct {
	app        *appObject
	currentPos mouse.MousePos
	prevPos    mouse.MousePos
}

func (m *MouseEvent) Initialize(app *appObject) {
	m.currentPos = mouse.MousePos{Point: device.Point{0, 0}}
	m.prevPos = mouse.MousePos{Point: device.Point{0, 0}}
	m.app = app
}

func (m *MouseEvent) MouseProc() {

	width, height := m.app.Screen.Main.GetSize()
	m.currentPos.SetPoint(robotgo.GetMousePos())

	if m.currentPos.Equals(&m.prevPos.Point) == false {
		cx, cy := m.currentPos.GetPoint()
		px, _ := m.prevPos.GetPoint()
		//event
		// 좌 끝단 도달
		if m.app.IsFoucs == true {
			if cx <= 0 && px <= 0 {
				fmt.Printf("좌 끝단 : x: %d y:  %d\n", cx, cy)
			} else if cx >= width-1 && px >= width-1 {
				ox, oy := m.currentPos.GetVelocity(m.prevPos)
				fmt.Printf("우 끝단 : x: %d y:  %d\n", ox, oy)
				m.app.IsFoucs = false
				focusChange := common.Message{}
				focusChange.Type = common.MSG_SCREEN
				focusChange.Code = common.SCREEN_FOCUS_CHANGE
				m.app.SendMessage <- focusChange
			}
			m.prevPos.Initialize(m.currentPos.Point)
		} else {
			ox, oy := m.currentPos.GetVelocity(m.prevPos)
			mouseEvent := mouse.Mouse{device.Point{cx, cy}, 2, 2}
			mouseMove := common.Message{}
			mouseMove.Type = common.MSG_MOUSE
			mouseMove.Code = common.MOUSE_MOVE
			mouseMove.IsServer = m.app.IsServer
			mouseMove.Message = common.ObjectToByte(mouseEvent)
			m.app.SendMessage <- mouseMove

			robotgo.MoveMouse(width-10, height-10)
			m.prevPos.X = width - 10 - ox
			m.prevPos.Y = height - 10 - oy
			m.prevPos.Initialize(m.currentPos.Point)
		}

	}
	// btn := robotgo.AddEvent("mleft")
	// if btn == 0 {
	// 	fmt.Printf("%d - %d\n", m.currentPos.X, m.prevPos.X)
	// }
}