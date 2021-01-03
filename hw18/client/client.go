package main

import (
	"bufio"
	"flag"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/websocket"
)

const (
	promt = "-> "
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", "localhost:8000", "http service address")

	var passw string
	flag.StringVar(&passw, "passw", "password", "service password")

	flag.Parse()

	usend := url.URL{Scheme: "ws", Host: addr, Path: "/send"}

	urecieve := url.URL{Scheme: "ws", Host: addr, Path: "/messages"}
	log.Printf("connecting to %s", urecieve.String())

	crecieve, _, err := websocket.DefaultDialer.Dial(urecieve.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer crecieve.Close()

	go func() {
		for {
			_, message, err := crecieve.ReadMessage()
			if err != nil {
				log.Println("error:", err)
				return
			}
			log.Printf("recv: %s", message)
			log.Print(promt)
		}
	}()

	reader := bufio.NewReader(os.Stdin)

	for {
		log.Print(promt)

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("read message error: %v", err)
			break
		}
		text = strings.Replace(text, "\n", "", -1)

		csend, _, err := websocket.DefaultDialer.Dial(usend.String(), nil)
		if err != nil {
			log.Fatal("dial error:", err)
		}

		err = csend.WriteMessage(websocket.TextMessage, []byte(passw))
		if err != nil {
			log.Fatalf("auth error: %v", err)
			break
		}

		_, ok, err := csend.ReadMessage()
		if err != nil {
			log.Fatal("read message error:", err)
			return
		}
		if string(ok) != "OK" {
			log.Fatalf("connection error: server return %s, but expected OK", string(ok))
		}

		err = csend.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			log.Fatalf("write message error: %v", err)
			break
		}

		csend.Close()
	}
}
