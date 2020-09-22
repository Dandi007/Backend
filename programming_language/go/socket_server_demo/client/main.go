package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type message struct {
	Name   string
	Scores []int
}

func main() {
	socketAddress := "/tmp/server.sock"
	time := 1000
	for {
		conn, err := net.Dial(
			"unix",
			socketAddress,
		)
		if err != nil {
			fmt.Println(err)
			break
		}

		newMessage := &message{
			Name:   "test",
			Scores: []int{1, 2, 3, 4, 5},
		}
		messageBytes, _ := json.Marshal(newMessage)

		nBytes, err := conn.Write(messageBytes)
		fmt.Printf("%d bytes send to socket,err message:%v\n", nBytes, err)
		time--
		if time == 0 {
			break
		}
		conn.Close()
	}
}
