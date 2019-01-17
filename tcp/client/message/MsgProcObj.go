package message

import (
	"MDIIC/common"
	"net"
)

type MessageProc interface {
	RecvMessage(msg common.Message) bool
	SendMessage(conn net.Conn)
}
