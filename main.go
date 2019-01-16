package main

import (
	"MDIIC/tcp/client"
	"MDIIC/tcp/server"
	"fmt"
	"time"
)

func main() {

	var choose int = 0
	for {
		fmt.Printf("Choose Server(1) or Client(2) \n")
		fmt.Scanf("%d", &choose)
		if choose == 1 || choose == 2 {
			break
		}
	}
	if choose == 1 {
		go server.ServerMain()
	} else if choose == 2 {
		go client.ClientMain()
	}
	//	go controller.AppMain()
	//	go tcp.TcpMain()
	for {
		time.Sleep(time.Millisecond * 500)
	}
}
