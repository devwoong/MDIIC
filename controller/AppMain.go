package controller

import (
	"MDIIC/common"
	screen "MDIIC/device/screen"
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
	IsServer        bool
	IsFoucs         bool
}

func (app *appObject) AppMain(IsServer bool) {
	// robotgo.ScrollMouse(10, "up")
	// robotgo.MouseClick("left", true)
	// robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)

	app.IsServer = IsServer
	app.IsFoucs = IsServer
	app.Screen.Main.SetSize(robotgo.GetScreenSize())

	mouseEvent := MouseEvent{}
	mouseEvent.Initialize(app)

	for {
		mouseEvent.MouseProc()
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
