package client

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// RunWSClient runs websocket client
func RunWSClient(serverURL string, header http.Header) {
	c, _, err := websocket.DefaultDialer.Dial(serverURL, header)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	c.WriteMessage(websocket.TextMessage, []byte("message 1"))
	c.WriteMessage(websocket.TextMessage, []byte("message 2"))
	c.WriteMessage(websocket.TextMessage, []byte("message 3"))

	err = c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close:", err)
		return
	}
}
