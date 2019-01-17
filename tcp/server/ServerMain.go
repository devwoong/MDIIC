package server

import (
	"MDIIC/common"
	"MDIIC/controller"
	"MDIIC/tcp/server/protocol"
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

var server protocol.Server

func ServerMain() {
	service := "0.0.0.0:1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal errors  : %s", err.Error())
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, "host Ip : %s \t host port : %s\n", tcpAddr.IP, tcpAddr.Port)
	listener, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal errors : %s", err.Error())
		os.Exit(1)
	}

	server = protocol.Server{}
	server.Initialize()

	go controller.GetInstance().AppMain(true)
	connectLoop(listener)
}
func connectLoop(listener *(net.TCPListener)) {
	var i int = 0
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go clientHandle(conn, i)
		//go isAlivceClient()
		i++
	}
}
func clientHandle(conn net.Conn, id int) {
	defer conn.Close()
	client := protocol.Client{}
	client.CreateClient(conn, id)
	server.AddClient(&client)

	fmt.Printf("connect client id : %d, ip : %s \n", id, server.Clients[id].Conn.RemoteAddr())
	fmt.Printf("current client connet num : %d \n", len(server.Clients))

	go sendMessageHandle(id)
	go recvMessageHandle(id)
	messageRead(id)
}

func messageRead(id int) {
	for {
		r := bufio.NewReader(server.Clients[id].Conn)
		Data := make([]byte, 1024)
		n, err := r.Read(Data)
		//n, err := conn.Read(readPacket.Data)
		if err == io.EOF {
			exitMessage := common.Message{}
			exitMessage.Type = common.MSG_EXIT
			server.Clients[id].SendPacket <- exitMessage
			server.Clients[id].RecvPacket <- exitMessage
			fmt.Printf("client exit id : %d, ip : %s \n", id, server.Clients[id].Conn.RemoteAddr())
			server.DeleteClient(id)
			fmt.Printf("current client connet num : %d \n", len(server.Clients))
			return
		}
		if n != 0 {
			var message common.Message
			err = common.ByteToObject(Data, &message)
			server.Clients[id].SendPacket <- message
			fmt.Printf("read size : %d \n", n)
		}

	}
}

func recvMessageHandle(id int) {
EXITRECV:
	for {
		select {
		case message := <-server.Clients[id].SendPacket:
			switch message.Type {
			case common.MSG_EXIT:
				break EXITRECV
			case common.MSG_ALIVE:
				{
					Message := common.Message{}
					Message.Type = common.MSG_ALIVE
					server.Clients[id].RecvPacket <- Message
				}
			// case "&&ALIVE&&":
			// 	{
			// 		server.Clients[id].IsAlive = true
			// 	}
			case common.MSG_STRING:
				{
					fmt.Printf("client %d message : %s \n", id, string(message.Message))
				}
			default:
				fmt.Printf("client %d message : %s \n", id, message.Type)
			}
		}
	}
}
func sendMessageHandle(id int) bool {
EXIT:
	for {
		select {
		case message := <-server.Clients[id].RecvPacket:
			switch message.Type {
			case common.MSG_EXIT:
				break EXIT
			case common.MSG_ALIVE:
				_, err := server.Clients[id].Conn.Write(common.ObjectToByte(message))
				if err != nil {

				}
			}
		case appMessage := <-controller.GetInstance().SendMessage:
			for _, c := range server.Clients {

				c.SendPacket <- appMessage
			}
		}
	}
	return true
}
