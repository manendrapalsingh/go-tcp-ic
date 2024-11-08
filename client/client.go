package main

import (
	"fmt"
	"net"
)

func main() {

	// to connect to the tcp server

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Errorf("Failed to connect to the tcp server")
		return
	}

	defer conn.Close()

	data := []byte("Hi there !")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Errorf("Failed to write data ")
	}

}
