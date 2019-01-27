package client

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

var tickCount int = 0
var movePos device.Point = device.Point{0, 0}

const MOVE_TICK int = 2

func (m *MouseEvent) MouseProc() {
	m.currentPos.SetPoint(robotgo.GetMousePos())
	cx, cy := m.currentPos.GetPoint()

	if m.currentPos.Equals(&m.prevPos.Point) == false && m.app.IsFoucs == true {
		px, _ := m.prevPos.GetPoint()
		//event
		// 좌 끝단 도달
		//if m.app.IsFoucs == true {
		if cx <= 0 && px <= 0 {
			fmt.Printf("좌 끝단 : x: %d y:  %d\n", cx, cy)
			m.app.IsFoucs = false
			focusChange := common.Message{}
			focusChange.Type = common.MSG_SCREEN
			focusChange.Code = common.SCREEN_FOCUS_LEFT_CHANGE
			mouseMsg := mouse.Mouse{}
			mouseMsg.X = cx
			mouseMsg.Y = cy
			focusChange.Message = common.ObjectToByte(mouseMsg)
			m.app.SendMessage <- focusChange
		}
		m.prevPos.Initialize(m.currentPos.Point)
	}
}

func (m *MouseEvent) MouseMove(x, y int) {
	robotgo.MoveMouse(m.currentPos.X+x, m.currentPos.Y+y)
}

func (m *MouseEvent) SetMousePos(x, y int) {
	robotgo.MoveMouse(x, y)
}
