package utils

import (
	"fmt"
	"net"
	"os"
)

func ParseIP(name string) {
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid address")
	} else {
		fmt.Println("The address is ", addr.String())
	}
	os.Exit(0)
}
