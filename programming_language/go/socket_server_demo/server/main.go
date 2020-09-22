package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	netAddress := "/tmp/server.sock"

	server, err := net.Listen(
		"unix",
		netAddress,
	)

	if err != nil {
		fmt.Println(err)
	}

	connCount := 0

	for {
		conn, _ := server.Accept()
		connCount++
		fmt.Printf("new conn incoming,total conn num:%d\n", connCount)
		go func(conn net.Conn) {
			data, err := ioutil.ReadAll(conn)
			if err != nil {
				// pass
			}
			fmt.Printf("message from client:\n%s\n", string(data))
		}(conn)
		if connCount > 10 {
			break
		}
	}
	server.Close()
}
