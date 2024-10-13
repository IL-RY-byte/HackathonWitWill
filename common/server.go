package common

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func StartServer() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("Server is running on :8080 (TLS)")
	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)) // TLS configuration
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	log.Println("Client connected")
	sessionKey, err := PerformKeyExchange()
	if err != nil {
		log.Println("Key exchange error:", err)
		return
	}

	err = conn.WriteJSON(sessionKey)
	if err != nil {
		log.Println("Failed to send session key:", err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received message: %s", string(message))
	}
}
