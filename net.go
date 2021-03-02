package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	nl, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	conn, err := nl.Accept()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	remotAddr := conn.RemoteAddr().String()
	fmt.Println(remotAddr)
	conn.Write([]byte("হা আমি করতে পেরেছি"))

	conn.Close()
	nl.Close()

}
