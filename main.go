package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/pkg/browser"
	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.Handle("/ws", websocket.Handler(handleWebSocket))

	go func() {
		if err := browser.OpenURL("http://localhost:3000"); err != nil {
			log.Fatal(err)
		}
	}()

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}

func handleWebSocket(ws *websocket.Conn) {
	defer ws.Close()

	if err := websocket.Message.Send(ws, "Hello! Please send messages."); err != nil {
		log.Fatal(err)
	}

	for {
		receiveMsg := ""
		err := websocket.Message.Receive(ws, &receiveMsg)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		sendMsg := fmt.Sprintf("received \"%s\"", receiveMsg)
		if err := websocket.Message.Send(ws, sendMsg); err != nil {
			log.Fatal(err)
		}
	}
}
