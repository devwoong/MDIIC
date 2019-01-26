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
	onEvent    bool
}

func (m *MouseEvent) Initialize(app *appObject) {
	m.currentPos = mouse.MousePos{Point: device.Point{0, 0}}
	m.prevPos = mouse.MousePos{Point: device.Point{0, 0}}
	m.app = app
	m.onEvent = false
}

func (m *MouseEvent) MouseProc() {

	width, height := m.app.Screen.Main.GetSize()
	pixWidth := width - 10
	pixHeight := height - 10
	m.currentPos.SetPoint(robotgo.GetMousePos())
	cx, cy := m.currentPos.GetPoint()

	if m.currentPos.Equals(&m.prevPos.Point) == false && m.app.IsFoucs == true {
		px, _ := m.prevPos.GetPoint()
		//event
		// 좌 끝단 도달
		//if m.app.IsFoucs == true {
		if cx <= 0 && px <= 0 {
			fmt.Printf("좌 끝단 : x: %d y:  %d\n", cx, cy)
			if m.app.IsServer == false {
				focusChange := common.Message{}
				focusChange.Type = common.MSG_SCREEN
				focusChange.Code = common.SCREEN_FOCUS_LEFT_CHANGE
				mouseMsg := mouse.Mouse{}
				mouseMsg.X = cx
				mouseMsg.Y = cy
				focusChange.Message = common.ObjectToByte(mouseMsg)
				m.app.SendMessage <- focusChange
			}
		} else if cx >= width-1 && px >= width-1 {
			ox, oy := m.currentPos.GetVelocity(m.prevPos)
			fmt.Printf("우 끝단 : x: %d y:  %d\n", ox, oy)
			if m.app.IsServer == true {
				m.app.IsFoucs = false
				focusChange := common.Message{}
				focusChange.Type = common.MSG_SCREEN
				focusChange.Code = common.SCREEN_FOCUS_RIGHT_CHANGE
				mouseMsg := mouse.Mouse{}
				mouseMsg.X = cx
				mouseMsg.Y = cy
				focusChange.Message = common.ObjectToByte(mouseMsg)
				m.app.SendMessage <- focusChange
			}
		}
		m.prevPos.Initialize(m.currentPos.Point)
	} else {
		// ox, oy := m.currentPos.GetVelocity(m.prevPos)
		if m.currentPos.X != pixWidth || m.currentPos.Y != pixHeight {
			fmt.Printf("우 : x: %d y:  %d\n", cx-pixWidth, cy-pixHeight)
			mouseEvent := mouse.Mouse{device.Point{cx, cy}, cx - pixWidth, cy - pixHeight}
			mouseMove := common.Message{}
			mouseMove.Type = common.MSG_MOUSE
			mouseMove.Code = common.MOUSE_MOVE
			mouseMove.IsServer = m.app.IsServer
			mouseMove.Message = common.ObjectToByte(mouseEvent)
			m.app.SendMessage <- mouseMove

			robotgo.MoveMouse(pixWidth, pixHeight)
		}
	}

	if m.onEvent == false {
		go m.onMouseEvent()
	}
}

func (m *MouseEvent) onMouseEvent() {
	m.onEvent = true
	btn := robotgo.AddEvent("mright")
	if btn == 0 {
		m.app.IsFoucs = true
		fmt.Printf("%d - %d\n", m.currentPos.X, m.prevPos.X)
	}
	m.onEvent = false
}

func (m *MouseEvent) MouseMove(x, y int) {
	robotgo.MoveMouse(m.currentPos.X+x, m.currentPos.Y+y)
}

func (m *MouseEvent) SetMousePos(x, y int) {
	robotgo.MoveMouse(x, y)
}
