package protocol

import (
	"MDIIC/common"
	"fmt"
	"net"
	"os"
	"sync"
)

type server struct {
	Service     string
	Conn        net.Conn
	RecvMessage chan common.Message
	SendMessage chan common.Message
}

func (s *server) ConnectServer() error {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", s.Service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	s.Conn = conn
	return err
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s\n", err.Error())
	}
}

var serverInstance *server
var mu sync.Mutex

func GetInstance() *server {
	mu.Lock()
	defer mu.Unlock()
	if serverInstance == nil {
		serverInstance = &server{"localhost:1201", nil, make(chan common.Message), make(chan common.Message)}
	}
	return serverInstance
}
