package message

import (
	"MDIIC/common"
	"MDIIC/tcp/client/protocol"
	"fmt"
	"net"
	"os"
	"time"
)

type AliveCheck struct {
	MessageProc
	ReConnectCnt int
	IsAlivce     bool
}

func (ac *AliveCheck) RecvMessage(msg common.Message) bool {
	fmt.Printf("isAlive \n")
	switch msg.Type {
	case common.MSG_ALIVE:
		ac.IsAlivce = true
	}
	return false
}

func (ac *AliveCheck) SendMessage(conn net.Conn) {
	for {
		sendMsg := common.Message{}
		sendMsg.Type = common.MSG_ALIVE
		conn.Write(common.ObjectToByte(sendMsg))

		time.Sleep(time.Second * 10)
		if ac.IsAlivce == false {
			fmt.Printf("re connecting.... %d \n", ac.ReConnectCnt)
			protocol.GetInstance().ConnectServer()
			ac.ReConnectCnt++
			if ac.ReConnectCnt >= 3 {
				fmt.Printf("re connect fail.... exit...\n")
				os.Exit(0)
			}
		} else {
			ac.IsAlivce = false
			ac.ReConnectCnt = 0
		}
	}
}
