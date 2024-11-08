package main

import (
	"fmt"
	"net"
)

func main() {

	// to listen the incoming request
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {

		fmt.Errorf("Error while creating tcp listener ", err)
		return

	}

	defer listener.Close()

	fmt.Println("tcp sever staring to listen on port 8080")

	// for accepting all the incoming request
	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Errorf("Failed to accept the tcp request", err)
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	// creating a buffer to read into
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Errorf("Failed to read from the buffer ", err)
			return
		}

		fmt.Println("Recived data /n", string(buffer[:n]))
	}

}
