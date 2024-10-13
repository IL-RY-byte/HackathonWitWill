package common

import (
	"bufio"
	"crypto/tls"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func StartClient() {
	tlsConfig := &tls.Config{InsecureSkipVerify: true}
	u := url.URL{Scheme: "wss", Host: "localhost:8080", Path: "/ws"}
	log.Printf("Connecting to %s", u.String())

	dialer := websocket.Dialer{TLSClientConfig: tlsConfig}
	c, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	defer c.Close()

	// Receive session key from the server
	var sessionKey []byte
	err = c.ReadJSON(&sessionKey)
	if err != nil {
		log.Fatal("Failed to receive session key:", err)
	}
	log.Printf("Received session key size: %d bytes", len(sessionKey))

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			var msg Message
			err := c.ReadJSON(&msg)
			if err != nil {
				log.Println("Read error:", err)
				return
			}
			decryptedContent, err := DecryptMessage(sessionKey, []byte(msg.Content))
			if err != nil {
				log.Println("Decrypt error:", err)
				continue
			}
			log.Printf("Received: %s from %s at %s", string(decryptedContent), msg.Sender, time.Unix(msg.Timestamp, 0))
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		log.Print("Enter message: ")
		if scanner.Scan() {
			message := scanner.Text()
			encryptedContent, err := EncryptMessage(sessionKey, []byte(message))
			if err != nil {
				log.Println("Encrypt error:", err)
				continue
			}
			msg := Message{
				Sender:    "Client",
				Content:   string(encryptedContent),
				Timestamp: time.Now().Unix(),
			}

			err = c.WriteJSON(msg)
			if err != nil {
				log.Println("Write error:", err)
				return
			}
		}
	}
}
