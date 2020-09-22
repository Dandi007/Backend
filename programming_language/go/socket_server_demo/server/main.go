package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

var netAddress = flag.String("server address", "/tmp/server.sock", "")
var serverConnType = flag.String("server connection type", "unix", "")

func main() {
	flag.Parse()

	server, err := net.Listen(
		*serverConnType,
		*netAddress,
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	connCount := 0

	for {
		conn, _ := server.Accept()
		connCount++
		fmt.Printf("new conn incoming,total conn num:%d\n", connCount)
		go func(conn net.Conn) {
			// 使用ioutil来避免多次read to buffer && 拼接成最终bytes
			// TODO:另一种方法的标准写法是?
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
