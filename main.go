package main

import (
	"networking/gui"
	"networking/protocols"
	"networking/utils"
)

func main() {
	protocols.DialTCP("hellow from client", "0.0.0.0", "8080")
	utils.ParseIP("127.0.0.1")
	gui.Window(800,600)
}

