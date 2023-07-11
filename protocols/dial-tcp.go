package protocols

import (
	"fmt"
	"net"
)

func DialTCP(message string, ip string, port string) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()


	fmt.Println("Connected to", conn.RemoteAddr())

	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	fmt.Println("Data sent to server:", message)
}

