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
		clientAtoB, err := net.DialTCP("tcp4", clientA, clientB)
		if err != nil {
			fmt.Println(err)
		} else {
			defer clientAtoB.Close()
			clientAtoB.SetLinger(0)
			clientAtoB.SetNoDelay(true)
			clientAtoB.SetKeepAlive(false)
			fmt.Println("connected as Client A!")
			buffer := make([]byte, 64)
			_, err = clientAtoB.Read(buffer)
			if err != nil {
				continue
			}
		}
		time.Sleep(time.Second)
	}
}
