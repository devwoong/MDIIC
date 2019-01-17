package client

import (
	"MDIIC/common"
	"MDIIC/controller"
	"MDIIC/tcp/client/message"
	"MDIIC/tcp/client/protocol"
	"bufio"
	"fmt"
	"time"
)

var messageProcs []message.MessageProc

func ClientMain() {
	// if len(os.Args) != 2 {
	// 	fmt.Println(os.Stderr, "Usage %s host:port", os.Args[0])
	// 	os.Exit(0)
	// }
	// service := os.Args[1]

	// service := "localhost:1201"

	for {
		server := ""
		fmt.Printf("Enter the Server address \n")
		fmt.Scanf("%s", &server)
		//server = "127.0.0.1"
		if server == "" {
			continue
		}
		err := protocol.GetInstance().ConnectServer()
		if err != nil {
			continue
		} else {
			break
		}
	}
	fmt.Printf("------starting client------")
	defer protocol.GetInstance().Conn.Close()

	messageProcs = []message.MessageProc{
		&message.AliveCheck{ReConnectCnt: 0, IsAlivce: true},
	}

	go recvMessage()
	go recvMsgProc()
	go tick()
	go sendMessage()
	go controller.GetInstance().AppMain(false)
	for {

	}
}
func recvMessage() {
	for {
		r := bufio.NewReader(protocol.GetInstance().Conn)
		Data := make([]byte, 1024)
		n, err := r.Read(Data)
		if err != nil {
			return
		}
		if n > 0 {
			var message common.Message
			common.ByteToObject(Data, &message)
			protocol.GetInstance().RecvMessage <- message
			// message := readPacket.UnPack(n)
		}
	}
}

func recvMsgProc() {
RECV_EXIT:
	for {
		select {
		case msg := <-protocol.GetInstance().RecvMessage:
			switch msg.Type {
			case common.MSG_EXIT:
				break RECV_EXIT
			case common.MSG_ALIVE:
				//log.Printf("Receive: %s \n", msg.Message)
				for _, v := range messageProcs {
					v.RecvMessage(msg)
				}
			}
		}
	}
}

func tick() {
	for _, v := range messageProcs {
		go v.SendMessage(protocol.GetInstance().Conn)
	}
	for {
		message := common.Message{}
		message.Type = common.MSG_STRING
		message.Message = []byte("TICK TICK")
		protocol.GetInstance().SendMessage <- message
		time.Sleep(time.Second * 1)
	}
}

func sendMessage() {
SEND_EXIT:
	for {
		select {
		case AppMsg := <-protocol.GetInstance().SendMessage:
			switch AppMsg.Type {
			case common.MSG_EXIT:
				break SEND_EXIT
			case common.MSG_ALIVE:
				protocol.GetInstance().Conn.Write(common.ObjectToByte(AppMsg))
			case common.MSG_STRING:
				protocol.GetInstance().Conn.Write(common.ObjectToByte(AppMsg))
				fmt.Printf(" :: %s\n", string(AppMsg.Message))
			}
		}
	}
}
