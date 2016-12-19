package main


import (
	"fmt"
	"./src/oxd-client/transport"
)

func main() {
	transport.SocketSend("oxd-server.com:8080","sad")
	fmt.Printf("hello, world\n")
}
