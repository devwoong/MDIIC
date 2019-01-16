package message

import (
	"MDIIC/tcp/client/protocol"
	"net"
)

type MessageProc interface {
	RecvMessage(msg protocol.Message) (bool, protocol.Message)
	SendMessage(conn net.Conn)
}
