package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	clientA, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf(":%v", "2222"))
	if err != nil {
		fmt.Println(err)
		return
	}

	clientB, err := net.ResolveTCPAddr("tcp4", fmt.Sprintf(":%v", "3333"))
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		clientBtoA, err := net.DialTCP("tcp4", clientB, clientA)
		if err != nil {
			fmt.Println(err)
		} else {
			defer clientBtoA.Close()
			clientBtoA.SetLinger(0)
			clientBtoA.SetNoDelay(true)
			clientBtoA.SetKeepAlive(false)
			fmt.Println("connected as Client B!")
			buffer := make([]byte, 64)
			_, err = clientBtoA.Read(buffer)
			if err != nil {
				continue
			}
		}
		time.Sleep(time.Second)
	}
}
