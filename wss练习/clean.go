package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

// https://blog.csdn.net/ALakers/article/details/111713405

func main() {
	url := "ws://localhost:8080/ws" // 服务器地址
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			err := ws.WriteMessage(websocket.BinaryMessage, []byte("ping"))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		_, data, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("receive: ", string(data))
	}
}
