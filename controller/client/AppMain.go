package client

import (
	"MDIIC/common"
	device "MDIIC/device/mouse"
	screen "MDIIC/device/screen"
	"bytes"
	"encoding/gob"
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
	IsFoucs         bool
}

var g_mouseEvent MouseEvent

func (app *appObject) AppMain() {
	app.IsFoucs = false
	app.Screen.Main.SetSize(robotgo.GetScreenSize())

	g_mouseEvent = MouseEvent{}
	g_mouseEvent.Initialize(app)

	go app.recvEvent()

	for {
		g_mouseEvent.MouseProc()
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

func (app *appObject) recvEvent() {
EVENTEXIT:
	for {
		select {
		case message := <-app.RecvMessage:
			switch message.Type {
			case common.MSG_EXIT:
				break EVENTEXIT
			case common.MSG_MOUSE:
				{
					if message.Code == common.MOUSE_MOVE {
						mouse := device.Mouse{}
						buf := bytes.NewBuffer(message.Message)
						d := gob.NewDecoder(buf)
						if err := d.Decode(&mouse); err != nil {
							panic(err)
						}
						g_mouseEvent.MouseMove(mouse.MoveX, mouse.MoveY)
						fmt.Printf("Move X : %d, Move Y : %d\n", mouse.MoveX, mouse.MoveY)
					}
				}
			case common.MSG_KEYBOARD:
				{

				}
			case common.MSG_SCREEN:
				{
					mouse := device.Mouse{}
					buf := bytes.NewBuffer(message.Message)
					d := gob.NewDecoder(buf)
					if err := d.Decode(&mouse); err != nil {
						panic(err)
					}

					switch message.Code {
					case common.SCREEN_FOCUS_RIGHT_CHANGE:
						app.IsFoucs = true
						g_mouseEvent.SetMousePos(0, mouse.Y)
						fmt.Printf("focus change true\n")
					}
				}
			default:
				{

				}
			}
		}
	}
}
