package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"log"
)

var origin = "http://localhost/"
var url = "ws://localhost:3000/ws"

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	message := []byte("{\"type\":\"broadcast\",\"payload\":20}")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	var msg = make([]byte, 512)
	_, err = ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg)
}
