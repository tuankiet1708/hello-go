package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			log.Println("Không thể nhận!")
			break
		}

		fmt.Println("Nhận từ client: " + reply)

		msg := "Đã nhận: " + reply
		fmt.Println("Gửi cho client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Không thể gửi!")
			break
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
